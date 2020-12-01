package app

import (
	"github.com/gin-gonic/gin"
)

// StartApp start sets up the server
func StartApp() {
	router := gin.Default()
	router.GET("/test")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
