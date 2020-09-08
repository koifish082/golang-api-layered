package search

import (
	"github.com/koifish082/golang-api-layered/src/app/application/dto"
	"github.com/koifish082/golang-api-layered/src/app/application/dto/search"
	"github.com/koifish082/golang-api-layered/src/app/domain/model"
)

func convertSearchRepositoryResponse(repositories *model.SearchRepository) (*search.RepositoryResponse, *dto.ClientError) {
	res := &search.RepositoryResponse{
	}
	return res, nil
}