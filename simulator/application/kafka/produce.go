package kafka

import (
	"encoding/json"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	routeStruct "github.com/phelipperibeiro/simulator/application/route"
	kafkaInfrastructure "github.com/phelipperibeiro/simulator/infra/kafka"
	"log"
	"os"
	"time"
)

// Produce is responsible to publish the positions of each request

// Example of a json request:
// {"clientId":"1","routeId":"1"}
// {"clientId":"2","routeId":"2"}
// {"clientId":"3","routeId":"3"}
// {"clientId":"3","routeId":"4"}

func Produce(msg *confluentKafka.Message) {
	producer := kafkaInfrastructure.NewKafkaProducer()
	route := routeStruct.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, position := range positions {
		kafkaInfrastructure.Publish(position, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 10000)
	}
}
