package v1

import (
	"net/http"

	"github.com/koifish082/golang-api-layered/src/app/interfaces/api/model/requestparam"

	"github.com/koifish082/golang-api-layered/src/app/application/usecase/search"

	"github.com/go-chi/chi"
	"github.com/koifish082/golang-api-layered/src/app/interfaces/api/middleware"
	"github.com/koifish082/golang-api-layered/src/app/interfaces/helper"
)

// SearchResources provide search interfaces
type SearchResources struct {
	Middleware               *middleware.Middleware
	GetRepositoriesInterface search.GetRepositoriesInterface
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
	r.With(rs.Middleware.ValidateSearchRepositoriesRequest).Get("/repositories", rs.getRepositories)

	return r
}

func (rs SearchResources) getRepositories(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	req := &search.GetRepositoriesRequest{
		Q: query.Get(requestparam.GetRepositoryQ),
	}
	rs.GetRepositoriesInterface = req
	res, err := rs.GetRepositoriesInterface.GetRepositories()
	if err != nil {
		helper.Fail(w, err.Code, err)
		return
	}
	helper.Succeed(w, res)
}
