package service

import (
	"fmt"
	"net/http/httptest"

	"github.com/koifish082/golang-api-layered/src/app/config"
	"github.com/koifish082/golang-api-layered/src/app/domain/model"
	"github.com/koifish082/golang-api-layered/src/app/library"
	"github.com/spf13/viper"
)

// SearchRequest common search parameter for search service function
type SearchRequest struct {
	MockServer *httptest.Server
}

// SearchInterface represent signatures of methods related to github search API.
type SearchInterface interface {
	SearchRepositories(q string) (*model.SearchRepository, *model.ServiceError)
}

// SearchRepositories search repositories on github
func (r *SearchRequest) SearchRepositories(q string) (*model.SearchRepository, *model.ServiceError) {
	endpoint := fmt.Sprintf("%s/search/repositories",
		viper.GetString(config.EnvGithubBaseURL),
	)
	headers := createCommonHeaders()
	queries := library.FilterNotEmptyValue(map[string]string{
		model.SearchRepositoryQueryQ: q,
	})

	// Activate mock server.
	if r.MockServer != nil {
		endpoint = r.MockServer.URL
	}

	resp, err := get(endpoint, headers, queries)
	if err != nil {
		return nil, err
	}

	res, e := model.UnmarshalSearchRepository(resp.Body())
	if e != nil {
		return nil, model.NewServiceError(e)
	}

	return &res, nil
}
