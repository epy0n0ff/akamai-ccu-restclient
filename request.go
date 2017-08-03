package restclient

// See:https://api.ccu.akamai.com/ccu/v2/docs/
type PurgeRequest struct {
	//Type string   `json:"omitempty`
	Objects []string `json:"objects"`
}
