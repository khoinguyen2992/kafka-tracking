package kafkatracking

import (
	"os"

	"github.com/khoinguyen2992/kafka-tracking/mongodb"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	KAFKA_TRACKING_COLLECTION = "kafka_tracking"
)

var db *mgo.Database

func init() {
	db = mongodb.Session.DB(os.Getenv("MONGODB_DATABASE"))
}

type KafkaTrackingRepository struct {
	Collection *mgo.Collection
}

func NewKafkaTrackingRepository() *KafkaTrackingRepository {
	return &KafkaTrackingRepository{
		Collection: db.C(KAFKA_TRACKING_COLLECTION),
	}
}

func (this *KafkaTrackingRepository) Create(kafkaTracking KafkaTracking) (KafkaTracking, error) {
	kafkaTracking.ID = bson.NewObjectId()
	err := this.Collection.Insert(kafkaTracking)
	return kafkaTracking, err
}

func (this *KafkaTrackingRepository) UpdateByRequestIDAtOP(requestID string, kafkaTracking KafkaTracking) (KafkaTracking, error) {
	err := this.Collection.Update(bson.M{
		"request_id": requestID,
	}, bson.M{
		"$set": bson.M{
			"arrive_op_at":  kafkaTracking.ArriveOPAt,
			"message_at_op": kafkaTracking.MessageAtOP,
		},
	})
	return kafkaTracking, err
}

func (this *KafkaTrackingRepository) UpdateByRequestIDAtNC(requestID string, kafkaTracking KafkaTracking) (KafkaTracking, error) {
	err := this.Collection.Update(bson.M{
		"request_id": requestID,
	}, bson.M{
		"$set": bson.M{
			"arrive_nc_at":  kafkaTracking.ArriveNCAt,
			"message_at_nc": kafkaTracking.MessageAtNC,
		},
	})
	return kafkaTracking, err
}
