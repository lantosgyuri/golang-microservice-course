package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/auth"
	"net/http"
)

// SignIn create a token for a valid user or returns an error response
func SignIn(c *gin.Context) {
	var req auth.Credentials

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid JSON body",
		})
		return
	}

	s := auth.Authenticator{
		DB: auth.UserProvider,
	}

	token, err := s.CreateToken(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			StatusCode: err.StatusCode,
			Message:    err.Message,
		})
		return
	}

	var t string
	t = string(*token)

	// This token should be set as a cookie
	c.JSON(http.StatusOK, struct {
		Token string
	}{
		Token: t,
	})
	return
}
