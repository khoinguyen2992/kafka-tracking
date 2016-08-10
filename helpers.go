package kafkatracking

import (
	"encoding/json"
	"time"
)

func GenerateRequestID() string {
	return time.Now().Format("20060102.150405.999999")
}

func ToJSON(object interface{}) string {
	data, _ := json.Marshal(object)
	return string(data)
}
