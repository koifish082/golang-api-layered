package model

import (
	"encoding/json"
	"net/http"
)

// UnmarshalServiceErrorResponse is convert bytes to ServiceError struct
func UnmarshalServiceErrorResponse(data []byte) (ServiceError, error) {
	var r ServiceError
	err := json.Unmarshal(data, &r)
	return r, err
}

// NewServiceError generate ServiceError struct
func NewServiceError(e error) *ServiceError {
	return &ServiceError{
		Code:    http.StatusInternalServerError,
		Message: e.Error(),
	}
}

// ServiceError response is error response from service layer
type ServiceError struct {
	Code    int
	Message string   `json:"message"`
	Errors  []Errors `json:"errors"`
}

// Errors struct for error detail
type Errors struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}
