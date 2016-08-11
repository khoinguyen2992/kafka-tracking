package kafkatracking

import (
	"encoding/json"
	"time"
)

type OrderProcessorConsumer struct {
}

func (this OrderProcessorConsumer) Serve(message string) {
	data := Message{}
	json.Unmarshal([]byte(message), &data)
	repo := NewKafkaTrackingRepository()
	now := time.Now()
	repo.UpdateByRequestIDAtOP(data.RequestID, KafkaTracking{
		ArriveOPAt:  &now,
		MessageAtOP: data.Message,
	})
	ProduceMessage(TOPIC_KAFKA_TRACKING_TO_NOTIFICATION_CENTER, message)
}

type NotificationCenterConsumer struct {
}

func (this NotificationCenterConsumer) Serve(message string) {
	data := Message{}
	json.Unmarshal([]byte(message), &data)
	repo := NewKafkaTrackingRepository()
	now := time.Now()
	repo.UpdateByRequestIDAtNC(data.RequestID, KafkaTracking{
		ArriveNCAt:  &now,
		MessageAtNC: data.Message,
	})
}
