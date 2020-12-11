package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator/provider/github"
	"net/http"
)

// CreateRepo handles the route to create one repository
func CreateRepo(c *gin.Context) {
	var request repositorycreator.RepoRequest
	provider := c.Param("provider")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid JSON body",
		})
		return
	}

	service, err := getRepoCreatorService(provider)
	if err != nil {
		createErrorResponse(c, err)
		return
	}

	resp, error := service.CreateSingleRepo(&request)

	if error != nil {
		c.JSON(error.StatusCode, ErrorResponse{
			StatusCode: error.StatusCode,
			Message:    error.Message,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateMultipleRepository handles the route to create multiple repositories
func CreateMultipleRepository(c *gin.Context) {
	var multipleRepoRequest repositorycreator.MultipleRepoRequest
	provider := c.Param("provider")

	if err := c.ShouldBindJSON(&multipleRepoRequest); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid JSON body",
		})
		return
	}

	service, err := getRepoCreatorService(provider)
	if err != nil {
		createErrorResponse(c, err)
		return
	}

	var requests []*repositorycreator.RepoRequest

	for i := 0; i < len(multipleRepoRequest.Repos); i++ {
		requests = append(requests, &multipleRepoRequest.Repos[i])
	}

	resp := service.CreateMutlipleRepos(requests)

	c.JSON(http.StatusOK, resp)
	fmt.Printf("Requesit is %v \n", resp)
}

func getRepoCreatorService(name string) (*repositorycreator.Service, error) {

	switch name {
	case "github":
		return &repositorycreator.Service{
			Provider: &github.Provider{},
		}, nil
	default:
		return nil, errors.New("There is no provider with given name")
	}
}

func createErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Message:    err.Error(),
	})
}
