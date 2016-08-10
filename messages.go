package kafkatracking

type Message struct {
	RequestID string `json:"request_id"`
	Message   string `json:"message"`
}
