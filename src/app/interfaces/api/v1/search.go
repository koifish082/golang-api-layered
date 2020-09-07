package v1

import (
	"net/http"

	"github.com/koifish082/golang-api-layered/src/app/application/usecase/search"

	"github.com/go-chi/chi"
	"github.com/koifish082/golang-api-layered/src/app/interfaces/api/middleware"
	"github.com/koifish082/golang-api-layered/src/app/interfaces/helper"
)

// SearchResources provide search interfaces
type SearchResources struct {
	Middleware       *middleware.Middleware
	GetCodeInterface search.GetCodeInterface
}

// NewSearchResources returns new SearchResources struct
func NewSearchResources(m *middleware.Middleware) *SearchResources {
	return &SearchResources{
		Middleware: m,
	}
}

// Routes creates a REST router for the search resource
func (rs SearchResources) Routes() chi.Router {
	r := chi.NewRouter()
	r.With(rs.Middleware.ValidateSearchCodeRequest).Get("/code", rs.getCode)

	return r
}

func (rs SearchResources) getCode(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	request := &search.GetCodeRequest{
		Q: query.Get("q"),
	}
	rs.GetCodeInterface = request
	res, err := rs.GetCodeInterface.GetCode()
	if err != nil {
		helper.Fail(w, err.Code, err)
		return
	}
	helper.Succeed(w, res)
}
