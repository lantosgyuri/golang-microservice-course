package auth

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

type mockedDb struct {
	get func(name string) (*User, error)
}

func (m *mockedDb) GetUser(userName string) (*User, error) {
	return m.get(userName)
}

// Only a permament test
func TestValidate(T *testing.T) {
	T.Run("Invalid token structure", func(t *testing.T) {

		mDb := mockedDb{
			get: func(name string) (*User, error) {
				return nil, nil
			},
		}

		s := Authenticator{
			DB: &mDb,
		}

		got := s.Validate("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCeyJVc2VybmFtZSI6IkRlbmVtIiwiZXhwIjoxNjA4NDY4NTAwfQ.nylOXs-1bi6oignCVKS5Mmz8hhuTQOGqSa4paRNxAPU")

		want := Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Is the token OK?",
		}

		assert.Equal(t, want.Message, got.Message)
		assert.Equal(t, want.StatusCode, got.StatusCode)
	})
}

func TestCreateToken(T *testing.T) {
	T.Run("user not exists", func(t *testing.T) {
		mDb := mockedDb{
			get: func(name string) (*User, error) {
				return nil, errors.New("User not found")
			},
		}

		s := Authenticator{
			DB: &mDb,
		}

		gotT, gotE := s.CreateToken(&Credentials{
			UserName: "t",
			Password: "t",
		})

		want := Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "User is not exists",
		}

		assert.NotNil(t, gotE)
		assert.Nil(t, gotT)

		assert.Equal(t, want.StatusCode, gotE.StatusCode)
		assert.Equal(t, want.Message, gotE.Message)
	})

	T.Run("password incorrect", func(t *testing.T) {
		mDb := mockedDb{
			get: func(name string) (*User, error) {
				return &User{
					ID:       1,
					UserName: "t",
					Password: "z",
				}, nil
			},
		}

		s := Authenticator{
			DB: &mDb,
		}

		gotT, gotE := s.CreateToken(&Credentials{
			UserName: "t",
			Password: "t",
		})

		want := Error{
			StatusCode: http.StatusUnauthorized,
			Message:    "Password is incorrect",
		}

		assert.NotNil(t, gotE)
		assert.Nil(t, gotT)

		assert.Equal(t, want.StatusCode, gotE.StatusCode)
		assert.Equal(t, want.Message, gotE.Message)

	})
}
