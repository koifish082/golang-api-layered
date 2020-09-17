package model

import (
	"encoding/json"
	"net/http"
)

// UnmarshalServerErrorResponse is convert bytes to ServiceError struct
func UnmarshalServerErrorResponse(data []byte) (ServerError, error) {
	var r ServerError
	err := json.Unmarshal(data, &r)
	return r, err
}

// NewServiceErrorByServerError generates Service error struct
func NewServiceErrorByServerError(servErr ServerError, code int) *ServiceError {
	return &ServiceError{
		Code:        code,
		ServerError: servErr,
	}
}

// NewServiceError generate ServiceError struct
func NewServiceError(e error) *ServiceError {
	return &ServiceError{
		Code: http.StatusInternalServerError,
		ServerError: ServerError{
			Message: e.Error(),
		},
	}
}

// ServiceError is error response from service layer
type ServiceError struct {
	Code int
	ServerError
}

// ServerError is response from External API
type ServerError struct {
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
}

// Error struct for error detail
type Error struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}
