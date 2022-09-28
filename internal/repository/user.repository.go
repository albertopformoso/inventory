package repository

import (
	"context"

	"github.com/albertopformoso/inventory/internal/entity"
)

const (
	queryInsertUser = `
	INSERT INTO user (email, name, password)
	VALUES (?, ?, ?);
	`
	queryGetUserByEmail = `
	SELECT id, email, name, password
	FROM user
	WHERE email = ?;
	`
)

func (r *repository) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, queryInsertUser, email, name, password)
	return err
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	user := &entity.User{}
	err := r.db.GetContext(ctx, user, queryGetUserByEmail, email)
	if err != nil {
		return nil, err
	}

	return user, err
}