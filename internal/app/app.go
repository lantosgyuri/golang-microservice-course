package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/create-repo/controllers"
)

// StartApp start sets up the server
func StartApp() {
	router := gin.Default()
	router.GET("/create", controllers.CreateRepo)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
