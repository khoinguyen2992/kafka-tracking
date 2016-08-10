package kafkatracking

import "os"

func ProduceMessage(topic, value string) bool {
	brokerList := os.Getenv("KAFKA_BROKER_LIST")
	kafkaClient := kafka.Init(brokerList)
	ok := kafkaClient.Producer.ProduceMessage(topic, value)
	return ok
}
