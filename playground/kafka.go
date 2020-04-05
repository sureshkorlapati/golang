package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	kafkaBroker := "foo:9092"
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBroker})
	if err != nil {
		fmt.Println("producer creation failed ", err.Error())
		return
	}

	topic := "bar"
	partition := kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}
	msg := &kafka.Message{TopicPartition: partition, Key: []byte("hello"), Value: []byte("world")}
	err = p.Produce(msg, nil)
	fmt.Println("done...")

	event := <-p.Events()
	switch e := event.(type) {
	case kafka.Error:
		pErr := e
		fmt.Println("producer error", pErr.String())
	default:
		fmt.Println("Kafka producer event", e)
	}
}
