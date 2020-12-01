package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateRepo handle the create route
func CreateRepo(c *gin.Context) {
	c.JSON(http.StatusOK, "It works")
}
