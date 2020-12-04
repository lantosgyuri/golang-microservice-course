package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator/controller"
)

// StartApp start sets up the server
func StartApp() {
	router := gin.Default()
	router.POST("/create/:provider", controller.CreateRepo)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
