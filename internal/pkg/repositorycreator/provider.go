package repositorycreator

// Provider interface describe the main functionalty of this package
type Provider interface {
	Create(request *RepoRequest) (*Repo, *RepoError)
}
