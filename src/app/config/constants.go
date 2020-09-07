package config

const (
	// Env is environment of the project
	Env = "ENV"
	// EnvPort is server port
	EnvPort = "PORT"
	// EnvLogLevel is min log level on logger
	EnvLogLevel = "LOG_LEVEL"
	// EnvGithubBaseURL is for github API
	EnvGithubBaseURL = "GITHUB_BASE_URL"
	// EnvGithubAPISecret is secret key for Github API
	EnvGithubAPISecret = "GITHUB_API_SECRET"
	// EnvAPIRequestTimeout is time out duration on API request
	EnvAPIRequestTimeout = "ENV_API_REQUEST_TIMEOUT"
	// EnvAPIRequestRetryCount is retry count on API request
	EnvAPIRequestRetryCount = "ENV_API_REQUEST_RESTY_COUNT"
)
