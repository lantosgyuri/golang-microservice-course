package repositorycreator

// RepoError is for describing which error ooccured during repo creation
type RepoError struct {
	StatusCode int
	Message    string
}
