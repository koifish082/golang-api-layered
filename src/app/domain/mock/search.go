package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/koifish082/golang-api-layered/src/app/domain/model"
)

const searchGetRepositoryJSONPath = "/search-get-repository.json"

// NewSearchRepositoriesMockServer creates new test server for Search repository.
func NewSearchRepositoriesMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Error conditions.
		q := r.URL.Query().Get(model.SearchRepositoryQueryQ)
		if q == "" {
			http.Error(w, newServerError("Missing required request parameter", nil), http.StatusBadRequest)
			return
		}

		// Read mock response.
		data, err := ioutil.ReadFile(getServiceMockDirectoryPath() + searchGetRepositoryJSONPath)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		fmt.Fprintln(w, string(data))
	}))
}

//SearchRequestMock struct
type SearchRequestMock struct {
}

// SearchRepositories is mock function used from usecase test
func (r *SearchRequestMock) SearchRepositories(q string) (*model.SearchRepository, *model.ServiceError) {
	// Mock error response.
	if q == "" {
		err := model.ServerError {
			Message: "Missing required request parameter",
		}
		return nil, model.NewServiceErrorByServerError(err, http.StatusBadRequest)
	}
	// Load mock response
	data, err := ioutil.ReadFile(getServiceMockDirectoryPath() + searchGetRepositoryJSONPath)
	if err != nil {
		return nil, model.NewServiceError(err)
	}
	res, err := model.UnmarshalSearchRepository(data)
	if err != nil {
		return nil, model.NewServiceError(err)
	}

	return res, nil
}

