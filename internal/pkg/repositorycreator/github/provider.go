package github

import (
	"github.com/lantosgyuri/golang-microservices-course/internal/pkg/repositorycreator"
)

// Service holds the repository creation logic
type Service struct{}

// Create created a repository on github
func (s *Service) Create(repoData *repositorycreator.Repo) (*repositorycreator.Repo, error) {
	return nil, nil
}
