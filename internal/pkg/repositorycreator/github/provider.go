package github

import (
	"encoding/json"
	"fmt"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator"
	"github.com/lantosgyuri/golang-microservices-course/pkg/restclient"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	authorazitationHeader = "Authorization"
)

// CreateRepoURL is the github url for creating new repositores
var CreateRepoURL = "https://https://api.guthub.com/user/repos"

// Service holds the repository creation logic
type Service struct{}

// Create created a repository on github
func (s *Service) Create(request *repositorycreator.RepoRequest) (*repositorycreator.Repo, *repositorycreator.RepoError) {
	headers := http.Header{}
	headers.Set(authorazitationHeader, fmt.Sprintf("token %s", os.Getenv("githubToken")))

	resp, err := restclient.Post(CreateRepoURL, request, headers)

	fmt.Println("start")
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
		fmt.Println("INSIDE 299")
		if err := json.Unmarshal(bytes, &errResponse); err != nil {
			fmt.Println("inside failed amrshal")
			return nil, &repositorycreator.RepoError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to umarshal the error response json body",
			}
		}

		fmt.Printf("marshal was ok %v", errResponse)
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
