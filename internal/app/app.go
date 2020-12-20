package app

import (
	"github.com/gin-gonic/gin"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/auth/handler"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/auth/middleware"
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator/controller"
	"github.com/lantosgyuri/golang-microservices-course/pkg/log"
	"net/http"
)

// StartApp start sets up the server
func StartApp() {
	log.InitLogger()
	router := gin.Default()
	router.Use(middleware.Authenticator)
	router.POST("/create/:provider", controller.CreateRepo)
	router.POST("/createMultiple/:provider", controller.CreateMultipleRepository)
	router.POST("/singin", handler.SignIn)
	router.GET("/hello", hello)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}

}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, struct {
		m string
	}{
		m: "Hello",
	})
}
