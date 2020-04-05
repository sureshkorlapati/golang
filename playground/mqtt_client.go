package main

import (
	"fmt"
	"log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	var opts = MQTT.NewClientOptions()
	opts.AddBroker("tcp://mqtt.eclipse.org:1883")
	opts.SetClientID("go-controller")

	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		log.Printf("topic: %s\n", msg.Topic())
	})

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Panicln(token.Error())
	}

	client.Subscribe("P1/#", 0, logHandler)
	// client.Subscribe("P1/controller/registration", 0, registrationHandler)

	for i := 0; i < 5; i++ {
		text := fmt.Sprintf("this is msg #%d!", i)
		token := client.Publish("P1/test", 0, false, text)
		token.Wait()
	}

	// time.Sleep(3 * time.Second)
}

var logHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	log.Printf("Topic %s logged...\n", msg.Topic())
}

var registrationHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	// log.Printf("Topic %s registered.\n", msg)
}
