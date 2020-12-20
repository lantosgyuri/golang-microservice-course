package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/auth"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/auth/handler"
	"net/http"
	"strings"
)

// Authenticator is a middleware for validating users
func Authenticator(c *gin.Context) {

	s := auth.Authenticator{
		DB: auth.UserProvider,
	}

	t := c.Request.Header["Token"]

	if err := s.Validate(auth.Token(strings.Join(t, ""))); err != nil {
		c.JSON(http.StatusBadRequest, handler.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid token Bruder",
		})
		return
	}
	c.Next()
}
