package restclient

import "fmt"

// DeleteResponse contains purge request results
type DeleteResponse struct {
	HttpStatus       int    `json:"httpStatus"`
	Detail           string `json:"detail"`
	EstimatedSeconds int    `json:"estimatedSeconds"`
	PurgeID          string `json:"purgeId"`
	SupportID        string `json:"supportId"`
}

// RateLimitResponse contains too many requests error response
type RateLimitResponse struct {
	HTTPStatus                  int    `json:"status"`
	RateLimitRemaining          int    `json:"rateLimitRemaining"`
	SupportID                   string `json:"supportId"`
	RateLimit                   int    `json:"rateLimit"`
	Title                       string `json:"title"`
	RateLimitCurrentRequestSize int    `json:"rateLimitCurrentRequestSize"`
}

// Error function implemented error interface
func (r RateLimitResponse) Error() string {
	return fmt.Sprintf("%d %s", r.HTTPStatus, r.Title)
}
