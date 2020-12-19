package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/auth/handler"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator/controller"
	"github.com/lantosgyuri/golang-microservices-course/pkg/log"
)

// StartApp start sets up the server
func StartApp() {
	log.InitLogger()
	router := gin.Default()
	router.POST("/create/:provider", controller.CreateRepo)
	router.POST("/createMultiple/:provider", controller.CreateMultipleRepository)
	router.POST("/singin", handler.SignIn)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}
