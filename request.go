package restclient

// See: https://developer.akamai.com/api/core_features/fast_purge/v3.html
type PurgeRequest struct {
	Objects []string `json:"objects"`
}
