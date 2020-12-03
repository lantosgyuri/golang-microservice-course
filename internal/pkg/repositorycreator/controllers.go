package repositorycreator

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateRepo handle the create route
func CreateRepo(c *gin.Context) {
	// check for which host and create this Repo, and call the interface method
	//validate json body, check if this is a valid Repo struct
	// calls provider
	// responds with an Repo struct what was created
	c.JSON(http.StatusOK, "It works")
}
