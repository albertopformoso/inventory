package service

import (
	"context"

	"github.com/albertopformoso/inventory/internal/model"
	"github.com/albertopformoso/inventory/internal/repository"
)

// Service is the business logic of the application.
//
//go:generate mockery --name=User --output=service --inpackage
type User interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)
}

type service struct {
	repository repository.User
}

func New(repository repository.User) User {
	return &service{repository: repository}
}
