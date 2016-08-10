package kakfatracking

import (
	"os"

	kafka "github.com/TranDuyThanh/kafka-client"
)

func ProduceMessage(topic, value string) bool {
	brokerList := os.Getenv("KAFKA_BROKER_LIST")
	kafkaClient := kafka.Init(brokerList)
	ok := kafkaClient.Producer.ProduceMessage(topic, value)
	return ok
}
