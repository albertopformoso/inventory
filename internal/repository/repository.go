package repository

import (
	"context"

	"github.com/albertopformoso/inventory/internal/entity"
	"github.com/albertopformoso/inventory/logger"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

// Repository is the interface that wraps the basic CRUD operations
//
//go:generate mockery --name=Repository --output=repository --inpackage
type Repository interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)

	SaveUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error
	GetUserRole(ctx context.Context, userID int64) ([]entity.UserRole, error)

	SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) error
	GetProducts(ctx context.Context) ([]entity.Product, error)
	GetProduct(ctx context.Context, id int64) (*entity.Product, error)
}

type repository struct {
	db  *sqlx.DB
	log *zerolog.Logger
}

func New(db *sqlx.DB) Repository {
	return &repository{
		db:  db,
		log: logger.New(),
	}
}
