package kafka

import (
	"fmt"
	"log"
	"os"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// KafkaConsumer holds all consumer logic and settings of Apache Kafka connections/
// Also has a Message channel which is a channel, where the messages are going to be pushed
type KafkaConsumer struct {
	MsgChan chan *confluentKafka.Message
}

// NewKafkaConsumer creates a new KafkaConsumer struct with its message channel as dependency
func NewKafkaConsumer(msgChan chan *confluentKafka.Message) *KafkaConsumer {
	return &KafkaConsumer{
		MsgChan: msgChan,
	}
}

// Consume consumes all message pulled from apache kafka and sent it to message channel
func (k *KafkaConsumer) Consume() {
	configMap := &confluentKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
		//"security.protocol": os.Getenv("security.protocol"),
		//"sasl.mechanisms":   os.Getenv("sasl.mechanisms"),
		//"sasl.username":     os.Getenv("sasl.username"),
		//"sasl.password":     os.Getenv("sasl.password"),
	}
	kafkaConsumer, err := confluentKafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}
	topics := []string{os.Getenv("KafkaReadTopic")}
	kafkaConsumer.SubscribeTopics(topics, nil)
	fmt.Println("Kafka consumer has been started")
	for {
		msg, err := kafkaConsumer.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}
