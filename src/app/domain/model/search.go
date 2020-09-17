package model

import (
	"encoding/json"
)

// SearchRepositoryQueryQ is a request parameter stands for search keyword
const SearchRepositoryQueryQ = "q"

// UnmarshalSearchRepository is convert bytes to UnmarshalSearchRepository struct
func UnmarshalSearchRepository(data []byte) (*SearchRepository, error) {
	var r SearchRepository
	err := json.Unmarshal(data, &r)
	return &r, err
}

// SearchRepository struct
type SearchRepository struct {
	TotalCount int    `json:"total_count"`
	Items      []Item `json:"items"`
}

// Item struct
type Item struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    Owner  `json:"owner"`
	HTMLURL  string `json:"html_url"`
}

// Owner struct
type Owner struct {
	Login string `json:"login"`
}
