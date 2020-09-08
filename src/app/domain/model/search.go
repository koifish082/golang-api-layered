package model

import (
	"encoding/json"
	"time"
)

// UnmarshalSearchRepository is convert bytes to UnmarshalSearchRepository struct
func UnmarshalSearchRepository(data []byte) (SearchRepository, error) {
	var r SearchRepository
	err := json.Unmarshal(data, &r)
	return r, err
}

// SearchRepository struct
type SearchRepository struct {
	TotalCount        int     `json:"total_count"`
	IncompleteResults bool    `json:"incomplete_results"`
	Items             []Items `json:"items"`
}

// Owner struct
type Owner struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
}

// Items struct
type Items struct {
	ID              int       `json:"id"`
	NodeID          string    `json:"node_id"`
	Name            string    `json:"name"`
	FullName        string    `json:"full_name"`
	Owner           Owner     `json:"owner"`
	Private         bool      `json:"private"`
	HTMLURL         string    `json:"html_url"`
	Description     string    `json:"description"`
	Fork            bool      `json:"fork"`
	URL             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	Homepage        string    `json:"homepage"`
	Size            int       `json:"size"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Language        string    `json:"language"`
	ForksCount      int       `json:"forks_count"`
	OpenIssuesCount int       `json:"open_issues_count"`
	MasterBranch    string    `json:"master_branch"`
	DefaultBranch   string    `json:"default_branch"`
	Score           float64   `json:"score"`
}
