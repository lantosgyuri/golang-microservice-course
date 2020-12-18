package handler

import (
	"cmd/go/internal/auth"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InMemoryDB jsut a fake DB
type InMemoryDB map[string]auth.User

var fakeDB = InMemoryDB{
    "user1": 
}

// DBProvider is a fek db connection
type DBProvider struct {
	db InMemoryDB
}

// GetUser returns a User from the DB
func (d *DBProvider) GetUser(userName string) (*User, error) {
	user, ok := db[userName]

	if !ok {
		return nil, errors.New("No user found")
	}

	return &user, nil
}

// SignIn validates the user
func SignIn(c *gin.Context) {
	var req auth.Credentials

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Statuscode: http.StatusBadRequest,
			Message:    "Invalid JSON body",
		})
		return
	}

	s := auth.Authenticator{
        DB: DBProvider{
            db: 
        }
    }

}
