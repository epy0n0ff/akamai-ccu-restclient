package restclient

// Response contains purge request results
type Response struct {
	HttpStatus       int    `json:"httpStatus"`
	Detail           string `json:"detail"`
	EstimatedSeconds int    `json:"estimatedSeconds"`
	PurgeId          string `json:"purgeId"`
	ProgressUri      string `json:"progressUuri"`
	PingAfterSeconds int    `json:"pingAfterSeconds"`
	SupportId        string `json:"supportId"`
}
