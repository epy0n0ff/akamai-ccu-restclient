package restclient

import "fmt"

// DeleteResponse contains purge request results
type DeleteResponse struct {
	HTTPStatus       int    `json:"httpStatus"`
	Detail           string `json:"detail"`
	EstimatedSeconds int    `json:"estimatedSeconds"`
	PurgeID          string `json:"purgeId"`
	SupportID        string `json:"supportId"`
}

// RateLimitResponse contains too many requests error response
type RateLimitResponse struct {
	HTTPStatus                  int    `json:"status"`
	Title                       string `json:"title"`
	RateLimitRemaining          int    `json:"rateLimitRemaining"`
	SupportID                   string `json:"supportId"`
	RateLimit                   int    `json:"rateLimit"`
	RateLimitCurrentRequestSize int    `json:"rateLimitCurrentRequestSize"`
}

// Error function implemented error interface
func (r RateLimitResponse) Error() string {
	return fmt.Sprintf("%d %s", r.HTTPStatus, r.Title)
}

// ErrorResponse contains general error response
type ErrorResponse struct {
	HTTPStatus  int    `json:"status"`
	Title       string `json:"title"`
	Detail      string `json:"detail"`
	DescribedBy string `json:"describedBy"`
}

// Error function implemented error interface
func (e ErrorResponse) Error() string {
	return fmt.Sprintf("%d %s", e.HTTPStatus, e.Title)
}
