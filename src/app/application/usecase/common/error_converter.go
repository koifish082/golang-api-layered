package common

import (
	"fmt"

	"github.com/koifish082/golang-api-layered/src/app/application/dto"

	"github.com/koifish082/golang-api-layered/src/app/domain/model"
)

// NewClientError convert from ServiceError to ClientError
func NewClientError(err *model.ServiceError) *dto.ClientError {
	var details []string
	for _, e := range err.Errors {
		details = append(details, fmt.Sprintf("%v", e))
	}
	return &dto.ClientError{
		Code:       err.Code,
		DetailCode: "",
		Message:    err.Message,
		Details:    details,
	}
}
