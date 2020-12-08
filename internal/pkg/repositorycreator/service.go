package repositorycreator

import (
	"fmt"
)

// Creator holds the logic what this package can do for you
type Creator interface {
	CreateSingleRepo(request *RepoRequest) (*Repo, *RepoError)
	CreateMutlipleRepos(multipleRequests []*RepoRequest) *MultitpleRepoResponse
}

// Service implements the Creator interface
type Service struct {
	Provider Provider
}

// CreateSingleRepo creates a single repository
func (s *Service) CreateSingleRepo(request *RepoRequest) (*Repo, *RepoError) {
	return s.Provider.Create(request)
}

// CreateMutlipleRepos creates a single repository
func (s *Service) CreateMutlipleRepos(multipleRequests []*RepoRequest) *MultitpleRepoResponse {

	resp, err := s.Provider.Create(&RepoRequest{
		Name:        "Name",
		Description: "Description",
		Private:     true,
	})

	var responses MultitpleRepoResponse

	fmt.Printf("RESPONSES %v", responses)

	if err != nil {
		responses.Errors = append(responses.Errors, *err)
	} else {
		responses.Repos = append(responses.Repos, *resp)
	}

	return &responses

}
