package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func create(c *gin.Context) {
	c.JSON(http.StatusOK, "It works")
}
