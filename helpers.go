package kafkatracking

import (
	"encoding/json"
	"fmt"
	"time"
)

func GenerateRequestID(extra string) string {
	return fmt.Sprintf("%s-%s", time.Now().Format("20060102.150405.999999999"), extra)
}

func ToJSON(object interface{}) string {
	data, _ := json.Marshal(object)
	return string(data)
}
