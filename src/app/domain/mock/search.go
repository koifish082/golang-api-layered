package mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/koifish082/golang-api-layered/src/app/domain/model"
)

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
		data, err := ioutil.ReadFile(getServiceMockDirectoryPath() + "/search-get-repository.json")
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		fmt.Fprintln(w, string(data))
	}))
}
