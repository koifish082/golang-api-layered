package search

// RepositoryResponse is client response on code
type RepositoryResponse struct {
	TotalCount   int
	Repositories []Repository
}

// Repository struct
type Repository struct {
	Name     string
	FullName string
	URL      string
	Owner    Owner
}

// Owner is repository's owner
type Owner struct {
	Name string
}
