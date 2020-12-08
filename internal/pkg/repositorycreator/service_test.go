package repositorycreator

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type createFunc = func(request *RepoRequest) (*Repo, *RepoError)
type testProvider struct {
	mockedCreate createFunc
}

func (t *testProvider) Create(request *RepoRequest) (*Repo, *RepoError) {
	return t.mockedCreate(request)
}

func TestCreateMultipleRepos(t *testing.T) {
	t.Run("No problem with single request", func(t *testing.T) {
		mCreate := func(request *RepoRequest) (*Repo, *RepoError) {
			return &Repo{
				Name:     request.Name,
				HomePage: "test.test.test",
				IsAdmin:  false,
			}, nil
		}
		p := testProvider{
			mockedCreate: mCreate,
		}
		s := Service{
			Provider: &p,
		}

		r := []*RepoRequest{
			&RepoRequest{
				Name:        "Name",
				Description: "Description",
				Private:     true,
			},
		}

		resp := s.CreateMutlipleRepos(r)

		want := MultitpleRepoResponse{
			Repos: []Repo{
				Repo{
					Name:     "Name",
					HomePage: "test.test.test",
					IsAdmin:  false,
				},
			},
		}

		assert.Equal(t, want.Repos, resp.Repos)
	})

	t.Run("Error with single request", func(t *testing.T) {
		mCreate := func(request *RepoRequest) (*Repo, *RepoError) {
			return nil, &RepoError{
				StatusCode: http.StatusBadRequest,
				Message:    "Something went wrong",
			}
		}
		p := testProvider{
			mockedCreate: mCreate,
		}
		s := Service{
			Provider: &p,
		}

		r := []*RepoRequest{
			&RepoRequest{
				Name:        "Name",
				Description: "Description",
				Private:     true,
			},
		}

		resp := s.CreateMutlipleRepos(r)

		want := MultitpleRepoResponse{
			Errors: []RepoError{
				RepoError{
					StatusCode: http.StatusBadRequest,
					Message:    "Something went wrong",
				},
			},
		}

		assert.Equal(t, want.Errors[0].StatusCode, resp.Errors[0].StatusCode)
		assert.Nil(t, resp.Repos)
	})
}
