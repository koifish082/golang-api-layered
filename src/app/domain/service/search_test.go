package service

import (
	"reflect"
	"testing"

	"github.com/koifish082/golang-api-layered/src/app/domain/mock"
	"github.com/koifish082/golang-api-layered/src/app/domain/model"
)

func TestSearchRepositories(t *testing.T) {

	// Request args.
	type args struct {
		q string
	}

	// All test cases.
	tests := []struct {
		name    string
		args    args
		wantRes *model.SearchRepository
		wantErr *model.ServiceError
	}{
		{
			name: "Search Repositories by search keyword (success)",
			args: args{
				q: "google",
			},
			wantRes: &model.SearchRepository{
				TotalCount: 40,
				Items: []model.Item{
					{
						Name:     "Tetris",
						FullName: "dtrupenn/Tetris",
						Owner: model.Owner{
							Login: "dtrupenn",
						},
						HTMLURL: "https://github.com/dtrupenn/Tetris",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "Search Repositories by search keyword (failed)",
			args: args{
				q: "",
			},
			wantRes: nil,
			wantErr: &model.ServiceError{
				Code: 400,
				ServerError: model.ServerError{
					Message: "Missing required request parameter",
				},
			},
		},
	}

	// Execute each test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// New request with mock server.
			request := &SearchRequest{
				MockServer: mock.NewSearchRepositoriesMockServer(),
			}
			defer request.MockServer.Close()

			// Call mock server through real service method.
			gotRes, gotErr := request.SearchRepositories(tt.args.q)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("SearchRepositories() got error = %v want error = %v", gotErr, tt.wantErr)
			}

			// Check service response is expected one.
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SearchRepositories() got response = %v want response = %v", gotRes, tt.wantRes)
			}
		})
	}
}
