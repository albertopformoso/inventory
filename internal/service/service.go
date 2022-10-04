package service

import (
	"context"

	"github.com/albertopformoso/inventory/internal/model"
	"github.com/albertopformoso/inventory/internal/repository"
	"github.com/albertopformoso/inventory/logger"
	"github.com/rs/zerolog"
)

// Service is the business logic of the application.
//
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*model.User, error)

	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error

	AddProduct(ctx context.Context, product model.Product, email string) error
	GetProducts(ctx context.Context) ([]model.Product, error)
	GetProduct(ctx context.Context, id int64) (*model.Product, error)
}

type service struct {
	repository repository.Repository
	log        *zerolog.Logger
}

func New(repository repository.Repository) Service {
	return &service{
		repository: repository,
		log:        logger.New(),
	}
}
