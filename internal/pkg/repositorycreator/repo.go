package repositorycreator

// Repo struct is the base type for a Repo
type Repo struct {
	Name     string
	HomePage string
	IsAdmin  bool
}

// RepoRequest holds the data of the repository what will be created
type RepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Private     bool   `json:"private"`
}

// MultitpleRepoResponse holds the created Repos or Errors what occured during creation
type MultitpleRepoResponse struct {
	Repos  []Repo
	Errors []RepoError
}
