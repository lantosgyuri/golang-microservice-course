package auth

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

// Token represents the JWT token
type Token string

// Service defines the auth logic
type Service interface {
	Validate(token Token) *Error
	CreateToken(credentials *Credentials) (*Token, *Error)
}

// Database defines the DB methods what needed for the AuthService
type Database interface {
	GetUser(username string) (*User, error)
}

// Authenticator authenticate the users
type Authenticator struct {
	DB Database
}

var jwtSecret = []byte("Should_ be_ a_secret")

// Validate validates the token
func (a *Authenticator) Validate(token Token) *Error {

	claims := &Claims{}
	var t string = string(token)

	tkn, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return &Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Is the token OK?",
		}
	}

	if !tkn.Valid {
		return &Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "Token is invalid",
		}
	}

	return nil
}

// CreateToken creates a token
func (a *Authenticator) CreateToken(credentials *Credentials) (*Token, *Error) {

	user, error := a.DB.GetUser(credentials.UserName)

	if error != nil {
		return nil, &Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "User is not exists",
		}
	}

	if user.Password != credentials.Password {
		return nil, &Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "Password is incorrect",
		}
	}

	experiationTime := time.Now().Add(5 * time.Minute)

	claims := Claims{
		Username: credentials.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: experiationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenToSend, err := token.SignedString(jwtSecret)

	if err != nil {
		return nil, &Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "Can not create token",
		}
	}

	var t Token = Token(tokenToSend)

	return &t, nil
}
