package main

import (
	"fmt"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	appProducer "github.com/phelipperibeiro/simulator/application/kafka"
	kafkaConsumer "github.com/phelipperibeiro/simulator/infra/kafka"

	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *confluentKafka.Message)
	consumer := kafkaConsumer.NewKafkaConsumer(msgChan)

	// new thread
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		// generate a new asynchronous thread
		go appProducer.Produce(msg)
	}
}

// Example of a json request:

// {"clientId":"1","routeId":"1"}
// {"clientId":"2","routeId":"2"}
// {"clientId":"3","routeId":"3"}

// ### to build kafka container
// terminal -> cd  .docker/kafka  -> docker-compose up -d

// to test (simulator) -> first terminal
// ### docker exec -it simulator bash
// 1) go run main.go

// to test (Kafka-producer) > second terminal -> send the message to -> new-direction's topic
// ### docker exec -it kafka_kafka_1 bash
// 2) kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction

// to test (Kafka-consumer)
// ### docker exec -it kafka_kafka_1 bash
// 3) kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal

//#########################################################################################################
//#########################################################################################################
//#########################################################################################################
//#########################################################################################################
//#########################################################################################################

// kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position
// kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal

// producer
// kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction

// consumer
// kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
