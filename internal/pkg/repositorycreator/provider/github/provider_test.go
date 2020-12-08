package github

import (
	"encoding/json"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	s := Provider{}
	t.Run("Test status >200", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			errorResponse := ErrorResponse{
				StatusCode:       http.StatusInternalServerError,
				Message:          "something went wrong",
				DocumantationURL: "",
				Errors:           []Error{},
			}
			bytes, _ := json.Marshal(errorResponse)
			w.Write(bytes)
		}))
		defer ts.Close()

		CreateRepoURL = ts.URL

		resp, err := s.Create(&repositorycreator.RepoRequest{
			Name:        "Name",
			Description: "Description",
			Private:     true,
		})

		want := repositorycreator.RepoError{
			StatusCode: http.StatusInternalServerError,
			Message:    "something went wrong",
		}
		assert.NotNil(t, err)
		assert.Equal(t, err.Message, want.Message)
		assert.Nil(t, resp)
	})

	t.Run("Test server not available", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			errorResponse := ErrorResponse{
				StatusCode:       http.StatusInternalServerError,
				Message:          "Request to external API failed",
				DocumantationURL: "",
				Errors:           []Error{},
			}
			bytes, _ := json.Marshal(errorResponse)
			w.Write(bytes)
		}))
		defer ts.Close()

		CreateRepoURL = ".d.d.d.d"

		resp, err := s.Create(&repositorycreator.RepoRequest{
			Name:        "Name",
			Description: "Description",
			Private:     true,
		})

		want := repositorycreator.RepoError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Request to external API failed",
		}

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, want.Message)
		assert.Equal(t, err.StatusCode, want.StatusCode)
		assert.Nil(t, resp)
	})

	t.Run("Test good reponse", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(http.StatusOK)
			response := CreateRepoResponse{
				ID:   1,
				Name: "Repo",
				Owner: RepoOwner{
					HTMLURL: "www.repo.repo",
				},
				Permission: RepoPermission{
					IsAdmin: true,
				},
			}
			bytes, _ := json.Marshal(response)
			w.Write(bytes)
		}))
		defer ts.Close()

		CreateRepoURL = ts.URL

		resp, err := s.Create(&repositorycreator.RepoRequest{
			Name:        "Repo",
			Description: "Description",
			Private:     true,
		})

		want := repositorycreator.Repo{
			Name:     "Repo",
			HomePage: "www.repo.repo",
			IsAdmin:  true,
		}

		assert.NotNil(t, resp)
		assert.Equal(t, want.Name, resp.Name)
		assert.Equal(t, want.IsAdmin, resp.IsAdmin)
		assert.Nil(t, err)
	})
}
