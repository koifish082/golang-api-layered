package search

import (
	"reflect"
	"testing"

	"github.com/koifish082/golang-api-layered/src/app/application/dto"
	"github.com/koifish082/golang-api-layered/src/app/application/dto/search"
	"github.com/koifish082/golang-api-layered/src/app/domain/mock"
)

func TestGetRepositories(t *testing.T) {

	// Args struct.
	type args struct {
		q string
	}

	// All test cases.
	tests := []struct {
		name    string
		args    *args
		wantRes *search.RepositoryResponse
		wantErr *dto.ClientError
	}{
		{
			name: "Search repository (success)",
			args: &args{
				q: "google",
			},
			wantRes: &search.RepositoryResponse{
				TotalCount: 40,
				Repositories: []search.Repository{
					{
						Name:     "Tetris",
						FullName: "dtrupenn/Tetris",
						URL:      "https://github.com/dtrupenn/Tetris",
						Owner: search.Owner{
							Name: "dtrupenn",
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "Search repositories with invalid arguments (failed)",
			args: &args{
				q: "",
			},
			wantRes: nil,
			wantErr: &dto.ClientError{
				Code:       400,
				DetailCode: "",
				Message:    "Missing required request parameter",
				Details:    []string{},
			},
		},
	}

	// Execute each test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// New request.
			r := &GetRepositoriesRequest{
				Q:                      tt.args.q,
				SearchServiceInterface: &mock.SearchRequestMock{},
			}
			// Call application method for getting mocked data.
			gotRes, gotErr := r.GetRepositories()
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("GetRepositories() got error = %v want error = %v", gotErr, tt.wantErr)
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetRepositories() got response = %v want response = %v", gotRes, tt.wantRes)
			}
		})
	}
}
