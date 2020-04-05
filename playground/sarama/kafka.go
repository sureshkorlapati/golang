package main

import (
	"log"

	kafka "github.com/Shopify/sarama"
)

func sendMessage() {
	producer, err := kafka.NewSyncProducer([]string{"52.183.6.88:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &kafka.ProducerMessage{Topic: "apache", Value: kafka.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}

func main() {
	sendMessage()
}
