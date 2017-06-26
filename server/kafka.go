package main

import (
	"github.com/Shopify/sarama"
  "fmt"
)
func main() {
  brokers := []string{"localhost:9202"}
  config := sarama.NewConfig()
  config.Producer.Partitioner = sarama.NewRandomPartitioner
  config.Producer.RequiredAcks = sarama.WaitForAll
  producer, _ := sarama.NewSyncProducer(brokers, config)

  topic := "test" //e.g create-user-topic
  partition := int32(1) //Wait for ll partitions to sync
  msg := "Hello Kafka" //e.g {"name":"John Doe", "email":"john.doe@email.com"}
  message := &sarama.ProducerMessage{
                   Topic: topic,
                   Partition: partition,
                   Value: sarama.StringEncoder(msg),
             }
  fmt.Println(message, producer)
  producer.SendMessage(message)
}
