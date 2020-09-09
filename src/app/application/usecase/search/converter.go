package search

import (
	"github.com/koifish082/golang-api-layered/src/app/application/dto/search"
	"github.com/koifish082/golang-api-layered/src/app/domain/model"
)

func convertSearchRepositoryResponse(repositoryResponse *model.SearchRepository) *search.RepositoryResponse {
	repositoryList := make([]search.Repository, 0, len(repositoryResponse.Items))
	for _, repo := range repositoryResponse.Items {
		r := search.Repository{
			Name:     repo.Name,
			FullName: repo.FullName,
			URL:      repo.URL,
			Owner: search.Owner{
				Name: repo.Owner.Login,
			},
		}
		repositoryList = append(repositoryList, r)
	}
	res := &search.RepositoryResponse{
		TotalCount:   repositoryResponse.TotalCount,
		Repositories: repositoryList,
	}

	return res
}
