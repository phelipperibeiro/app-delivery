package kafka

import (
	"log"
	"os"

	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewKafkaProducer creates a ready to go kafka.Producer instance
func NewKafkaProducer() *confluentKafka.Producer {
	configMap := &confluentKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"security.protocol": os.Getenv("security.protocol"),
		"sasl.mechanisms":   os.Getenv("sasl.mechanisms"),
		"sasl.username":     os.Getenv("sasl.username"),
		"sasl.password":     os.Getenv("sasl.password"),
	}
	kafkaProducer, err := confluentKafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return kafkaProducer
}

// Publish is simple function created to publish new message to kafka
func Publish(msg string, topic string, producer *confluentKafka.Producer) error {
	message := &confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{Topic: &topic, Partition: confluentKafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}
