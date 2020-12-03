package repositorycreator

// Service interface describe the main functionalty of this package
type Service interface {
	Create(name *RepoRequest) (*Repo, *RepoError)
}
