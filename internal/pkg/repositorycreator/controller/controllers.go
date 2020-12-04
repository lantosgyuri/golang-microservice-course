package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator/github"
	"net/http"
)

// CreateRepo handle the create route
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
		c.JSON(http.StatusBadRequest, ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
		})
		return
	}

	resp, error := service.Create(&request)

	if error != nil {
		c.JSON(error.StatusCode, ErrorResponse{
			StatusCode: error.StatusCode,
			Message:    error.Message,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func getRepoCreatorService(name string) (repositorycreator.Service, error) {

	switch name {
	case "github":
		return &github.Service{}, nil
	default:
		return nil, errors.New("There is no provider with given name")
	}
}
