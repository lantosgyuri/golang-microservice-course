package repositorycreator

// Repo struct is the base type for a Repo
type Repo struct {
	Name     string
	HomePage string
	IsAdmin  bool
}

// RepoRequest is the name of the repository what will be created
type RepoRequest struct {
	Name        string
	Description string
	Homepage    string
	Private     bool
}
