package github

// CreateRepoRequest is a modell for the request body
type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

// CreateRepoResponse is a modell for the response body
type CreateRepoResponse struct {
	ID         int64          `json:"id"`
	Name       string         `json:"name"`
	FullName   string         `json:"full_name"`
	Owner      RepoOwner      `json:"owner"`
	Permission RepoPermission `json:"permissions"`
}

// RepoOwner is a modell for the owner field in the createRepoResponse
type RepoOwner struct {
	ID      int64  `json:"id"`
	Login   string `json:"login"`
	URL     string `json:"url"`
	HTMLURL string `json:"html_url"`
}

// RepoPermission is a modell for the permission field in the createResponse
type RepoPermission struct {
	IsAdmin bool `json:"admin"`
	HasPull bool `json:"push"`
	HasPush bool `json:"pull"`
}

// ErrorResponse is a modell for the error response body
type ErrorResponse struct {
	StatusCode       int     `json:"status_code"`
	Message          string  `json:"message"`
	DocumantationURL string  `json:"documentation_url"`
	Errors           []Error `json:"errors"`
}

// Error is a modell for the Errors field in the ErrorResponse
type Error struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
