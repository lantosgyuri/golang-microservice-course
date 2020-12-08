package github

import (
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator"
	"github.com/lantosgyuri/golang-microservices-course/pkg/restclient"
	"io/ioutil"
	"net/http"
)

const (
	authorazitationHeader = "Authorization"
)

// CreateRepoURL is the github url for creating new repositores
var CreateRepoURL = "https://api.github.com/user/repos"

// Provider holds the repository creation logic
type Provider struct{}

// Create creates a repository on github
func (p *Provider) Create(request *repositorycreator.RepoRequest) (*repositorycreator.Repo, *repositorycreator.RepoError) {
	headers := http.Header{}
	headers.Set(authorazitationHeader, fmt.Sprintf("token %s", "ae5af3906b3f24b0cf7f09d9bc7844008b0abc5a"))

	resp, err := restclient.Post(CreateRepoURL, request, headers)
	if err != nil {
		return nil, &repositorycreator.RepoError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Request to external API failed",
		}
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		var errResponse ErrorResponse
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			return nil, &repositorycreator.RepoError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to umarshal the error response json body",
			}
		}

		return nil, &repositorycreator.RepoError{
			StatusCode: resp.StatusCode,
			Message:    errResponse.Message,
		}
	}

	var createRepoResp CreateRepoResponse

	if err := json.Unmarshal(bytes, &createRepoResp); err != nil {
		return nil, &repositorycreator.RepoError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to umarshal the response json body",
		}
	}

	return &repositorycreator.Repo{
		Name:     createRepoResp.Name,
		HomePage: createRepoResp.Owner.HTMLURL,
		IsAdmin:  createRepoResp.Permission.IsAdmin,
	}, nil
}
