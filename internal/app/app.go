package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator"
)

// StartApp start sets up the server
func StartApp() {
	router := gin.Default()
	router.GET("/create", repositorycreator.CreateRepo)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
