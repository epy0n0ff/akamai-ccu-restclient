package restclient

import "time"

// PurgeResponse contains purge request results
type PurgeResponse struct {
	Detail           string `json:"detail"`
	EstimatedSeconds int    `json:"estimatedSeconds"`
	HttpStatus       int    `json:"httpStatus"`
	PurgeId          string `json:"purgeId"`
	ProgressUri      string `json:"progressUri"`
	PingAfterSeconds int    `json:"pingAfterSeconds"`
	SupportId        string `json:"supportId"`
}

// QueueResponse contains queue check request results
type QueueResponse struct {
	Detail      string `json:"detail"`
	HttpStatus  int    `json:"httpStatus"`
	QueueLength int    `json:"queueLength"`
	SupportId   string `json:"supportId"`
}

// PurgeStatusResponse contains purge status check request results
type PurgeStatusResponse struct {
	CompletionTime           *time.Time `json:"completionTime"`
	HttpStatus               int        `json:"httpStatus"`
	OriginalEstimatedSeconds int        `json:"originalEstimatedSeconds"`
	OriginalQueueLength      int        `json:"originalQueueLength"`
	PingAfterSeconds         int        `json:"pingAfterSeconds"`
	ProgressUri              string     `json:"progressUri"`
	PurgeId                  string     `json:"purgeId"`
	PurgeStatus              string     `json:"purgeStatus"`
	SubmissionTime           *time.Time `json:"submissionTime"`
	SubmittedBy              string     `json:"submittedBy"`
	SupportId                string     `json:"supportId"`
}
