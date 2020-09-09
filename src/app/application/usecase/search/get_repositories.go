package search

import (
	"github.com/koifish082/golang-api-layered/src/app/application/dto"
	"github.com/koifish082/golang-api-layered/src/app/application/dto/search"
	"github.com/koifish082/golang-api-layered/src/app/application/usecase/common"
	"github.com/koifish082/golang-api-layered/src/app/domain/service"
)

// GetRepositoriesInterface interface
type GetRepositoriesInterface interface {
	GetRepositories() (*search.RepositoryResponse, *dto.ClientError)
}

// GetRepositoriesRequest is request struct for GetCode API
type GetRepositoriesRequest struct {
	Q                      string
	SearchServiceInterface service.SearchInterface
}

func (r *GetRepositoriesRequest) injectRepositoryServiceRequest() service.SearchInterface {
	if r.SearchServiceInterface != nil {
		return r.SearchServiceInterface
	}
	return &service.SearchRequest{}
}

// GetRepositories get Repositories
func (r *GetRepositoriesRequest) GetRepositories() (*search.RepositoryResponse, *dto.ClientError) {
	res, err := r.injectRepositoryServiceRequest().SearchRepositories(r.Q)
	if err != nil {
		return nil, common.NewClientError(err)
	}
	resp := convertSearchRepositoryResponse(res)
	return resp, nil
}
