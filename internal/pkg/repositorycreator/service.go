package repositorycreator

import (
	"sync"
)

type repoAndError struct {
	rep    *Repo
	repErr *RepoError
}

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

	var wg sync.WaitGroup
	respChan := make(chan *repoAndError)
	resultChan := make(chan *MultitpleRepoResponse)

	for _, req := range multipleRequests {
		wg.Add(1)
		go func(output chan *repoAndError, request *RepoRequest) {
			resp, err := s.CreateSingleRepo(request)

			respChan <- &repoAndError{
				rep:    resp,
				repErr: err,
			}

		}(respChan, req)
	}

	go s.handleResults(respChan, resultChan, &wg)

	wg.Wait()
	close(respChan)

	responses := <-resultChan
	return responses

}

func (s *Service) handleResults(respChan chan *repoAndError, result chan *MultitpleRepoResponse, wg *sync.WaitGroup) {
	var responses MultitpleRepoResponse

	for res := range respChan {
		switch {
		case res.rep != nil:
			responses.Repos = append(responses.Repos, *res.rep)
		case res.repErr != nil:
			responses.Errors = append(responses.Errors, *res.repErr)
		}
		wg.Done()
	}

	result <- &responses
}
