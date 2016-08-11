package kafkatracking

import (
	"fmt"
	"time"

	"github.com/khaiql/beego"
)

type APIGatewayHandler struct {
	beego.Controller
}

func (this *APIGatewayHandler) Get() {
	id := this.Ctx.Input.Param(":id")
	message := Message{
		RequestID: GenerateRequestID(id),
		Message:   fmt.Sprintf("%s-%s", TOPIC_KAFKA_TRACKING_TO_ORDER_PROCESSOR, id),
	}

	repo := NewKafkaTrackingRepository()
	now := time.Now()
	repo.Create(KafkaTracking{
		RequestID:   message.RequestID,
		ArriveAGAt:  &now,
		MessageAtAG: message.Message,
	})
	ProduceMessage(TOPIC_KAFKA_TRACKING_TO_ORDER_PROCESSOR, ToJSON(message))
}
