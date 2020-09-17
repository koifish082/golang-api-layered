package mock

import (
	"encoding/json"
	"os"

	"github.com/koifish082/golang-api-layered/src/app/domain/model"
)

func getServiceMockDirectoryPath() string {
	return os.Getenv("GOPATH") + "/src/github.com/koifish082/golang-api-layered/mock/service/"
}

// Generate new error for mock servers.
func newServerError(errMsg string, details []model.Error) string {
	errorResponse := model.ServerError{
		Message: errMsg,
		Errors:  details,
	}
	errRes, _ := json.Marshal(errorResponse)
	return string(errRes)
}
