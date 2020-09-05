package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/koifish082/golang-api-layered/src/app/interfaces/api/middleware"
	"github.com/koifish082/golang-api-layered/src/app/interfaces/helper"
)

type SearchResources struct {
	Middleware *middleware.Middleware
}

func NewSearchResources(m *middleware.Middleware) *SearchResources {
	return &SearchResources{
		Middleware: m,
	}
}

func (rs SearchResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.With(rs.Middleware.ValidateSearchCodeRequest).Get("/code", rs.getCode)

	return r
}

func (rs SearchResources) getCode(w http.ResponseWriter, r *http.Request) {
	helper.Succeed(w, "Hi!Code!")
}
