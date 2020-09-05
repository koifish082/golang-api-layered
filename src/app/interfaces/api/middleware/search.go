package middleware

import (
	"net/http"

	"github.com/koifish082/golang-api-layered/src/app/interfaces/helper"
)

func (m *Middleware) ValidateSearchCodeRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		if query.Get("q") == "" {
			helper.Fail(w, http.StatusBadRequest, "Invalid request parameter")
			return
		}
		next.ServeHTTP(w, r)
	})
}
