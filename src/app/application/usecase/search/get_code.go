package search

import (
	"github.com/koifish082/golang-api-layered/src/app/application/dto"
	"github.com/koifish082/golang-api-layered/src/app/application/dto/search"
	"github.com/koifish082/golang-api-layered/src/app/domain/service"
)

// GetCodeInterface interface
type GetCodeInterface interface {
	GetCode() (*search.CodeResponse, *dto.ClientError)
}

// GetCodeRequest is request struct for GetCode API
type GetCodeRequest struct {
	Q                      string
	SearchServiceInterface service.SearchInterface
}

func (r *GetCodeRequest) injectRepositoryServiceRequest() service.SearchInterface {
	if r.SearchServiceInterface != nil {
		return r.SearchServiceInterface
	}
	return &service.SearchRequest{}
}

// GetCode get Code
func (r *GetCodeRequest) GetCode() (*search.CodeResponse, *dto.ClientError) {
	r.injectRepositoryServiceRequest().SearchCode(r.Q)
	return nil, nil
}
