package service

import (
	"forum/internal/repo"
	"forum/models"
)

type service struct {
	repo repo.RepoI
}

type ServiceI interface {
	UserServiceI
}

type UserServiceI interface {
	GetUser(int) *models.User
	CreateUser(models.User) error
	Authenticate(string, string) (*models.Session, error)
	DeleteSession(string) error
}

func New(r repo.RepoI) ServiceI {
	return &service{
		r,
	}
}
