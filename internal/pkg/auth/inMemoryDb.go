package auth

import (
	"errors"
)

// InMemoryDb is faking a real DB
type InMemoryDb map[string]User

// DbProvider represents the connection to the DB
type DbProvider struct {
	db InMemoryDb
}

// GetUser returns a user from the DB or an error
func (d *DbProvider) GetUser(userName string) (*User, error) {
	user, ok := d.db[userName]

	if !ok {
		return nil, errors.New("No user found")
	}

	return &user, nil
}

// UserProvider is a fake DB for stroing users
var UserProvider *DbProvider = &DbProvider{
	db: InMemoryDb{
		"Denem": User{
			ID:       1,
			UserName: "Denem",
			Password: "Tom",
		},
	},
}
