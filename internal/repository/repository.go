package repository

import (
	"context"

	"github.com/albertopformoso/inventory/internal/entity"
	"github.com/jmoiron/sqlx"
)

// Repository is the interface that wraps the basic CRUD operations
//
//go:generate mockery --name=User --output=repository --inpackage
type User interface {
	SaveUser(ctx context.Context, email, name, password string) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) User {
	return &repository{db: db}
}
