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
	SearchCode(q string) (*model.SearchCode, *model.ServiceError)
}

// SearchCode search code on github
func (r *SearchRequest) SearchCode(q string) (*model.SearchCode, *model.ServiceError) {
	endpoint := fmt.Sprintf("%s/search/code",
		viper.GetString(config.EnvGithubBaseURL),
	)
	headers := createCommonHeaders()
	queries := library.FilterNotEmptyValue(map[string]string{
		"q": q,
	})

	// Activate mock server.
	if r.MockServer != nil {
		endpoint = r.MockServer.URL
	}

	resp, err := get(endpoint, headers, queries)
	if err != nil {
		return nil, err
	}

	res, e := model.UnmarshalSearchCode(resp.Body())
	if e != nil {
		return nil, model.NewServiceError(e)
	}

	return &res, nil
}
