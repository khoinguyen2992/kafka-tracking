package kafkatracking

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type KafkaTracking struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	RequestID   string        `json:"request_id" bson:"request_id"`
	ArriveAGAt  *time.Time    `json:"arrive_ag_at" bson:"arrive_ag_at"`
	MessageAtAG string        `json:"message_at_ag" bson:"message_at_ag"`
	ArriveOPAt  *time.Time    `json:"arrive_op_at" bson:"arrive_op_at"`
	MessageAtOP string        `json:"message_at_op" bson:"message_at_op"`
	ArriveNCAt  *time.Time    `json:"arrive_nc_at" bson:"arrive_nc_at"`
	MessageAtNC string        `json:"message_at_nc" bson:"message_at_nc"`
}
