package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// Credentials represents the info needed for auth
type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// Claims a truct that will be encoded to JWT
type Claims struct {
	Username string
	jwt.StandardClaims
}
