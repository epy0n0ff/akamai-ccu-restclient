package restclient

// PurgeResponse contains purge request results
type PurgeResponse struct {
	HttpStatus       int    `json:"httpStatus"`
	Detail           string `json:"detail"`
	EstimatedSeconds int    `json:"estimatedSeconds"`
	PurgeId          string `json:"purgeId"`
	ProgressUri      string `json:"progressUuri"`
	PingAfterSeconds int    `json:"pingAfterSeconds"`
	SupportId        string `json:"supportId"`
}

// QueueResponse contains queue check request results
type QueueResponse struct {
	HttpStatus  int    `json:"httpStatus"`
	QueueLength int    `json:"queueLength"`
	Detail      string `json:"detail"`
	SupportId   string `json:"supportId"`
}
