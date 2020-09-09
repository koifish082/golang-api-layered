package dto

// ClientError is for common error response for client.
type ClientError struct {
	Code       int      `json:"code"`
	DetailCode string   `json:"detailCode"`
	Message    string   `json:"message"`
	Details    []string `json:"details"`
}
