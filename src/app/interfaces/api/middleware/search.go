package middleware

import (
	"net/http"

	"github.com/koifish082/golang-api-layered/src/app/interfaces/api/model/requestparam"

	"github.com/koifish082/golang-api-layered/src/app/interfaces/helper"
)

// ValidateSearchRepositoriesRequest validate request parameter on search code API
func (m *Middleware) ValidateSearchRepositoriesRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if query.Get(requestparam.GetRepositoryQ) == "" {
			helper.Fail(w, http.StatusBadRequest, "Invalid request parameter")
			return
		}
		next.ServeHTTP(w, r)
	})
}
