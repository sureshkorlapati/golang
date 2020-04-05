package problem_reporter

import (
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/prometheus/common/log"
)

// var wg sync.WaitGroup
var mqttClient MQTT.Client

func main() {
	InitAndStart()

	fmt.Scanln()

}

func createMQTTClient(url string) (MQTT.Client, error) {
	opts := MQTT.NewClientOptions().AddBroker(url)
	opts.SetClientID("simba-client-" + string(time.Now().Unix()))
	// opts.SetDefaultPublishHandler(msgHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	log.Info("mqtt connection to " + url + " established ..")
	return client, nil
}

func Init() error {
	client, err := createMQTTClient("tcp://mqtt.eclipse.org:1883")
	if err != nil {
		log.Infof("error: %v", err)
		return err
	}
	mqttClient = client
	return nil
}

// make a channel
var messageChan = make(chan string)

func InitAndStart() error {

	err := Init()
	if err != nil {
		return err
	}

	go func() {
		mqttClient.Subscribe("topic/problem-reporter", 0, func(client MQTT.Client, msg MQTT.Message) {
			log.Info("subscribe called ...")
			messageChan <- string(msg.Payload())
			// GeneratePRT(string(msg.Payload()))

		})
	}()
	go handleMessage()

	return nil
}

func handleMessage() {

	for {
		select {
		case msg := <-messageChan:
			go GeneratePRT(msg)
		case <-time.After(time.Minute):
			fmt.Println("no messages in 1 min.")
		}

	}

}

func GeneratePRT(msg string) {
	log.Info("Generating PRT : " + msg)
}
