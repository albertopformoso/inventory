package repository

import (
	"context"

	"github.com/albertopformoso/inventory/internal/entity"
)

const (
	queryInsertUser = `
		INSERT INTO user (email, name, password)
		VALUES (?, ?, ?);`
	queryGetUserByEmail = `
		SELECT id, email, name, password
		FROM user
		WHERE email = ?;`

	queryInsertUserRole = `
		INSERT INTO user_role (user_id, role_id) VALUES(:user_role, :role_id);`
	queryRemoveUserRole = `
		DELETE FROM user_role WHERE user_id = :user_id and role_id = :role_id;`
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

func (r *repository) SaveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, queryInsertUserRole, data)

	return err
}

func (r *repository) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	data := entity.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err := r.db.NamedExecContext(ctx, queryRemoveUserRole, data)

	return err
}

func (r *repository) GetUserRole(ctx context.Context, userID int64) ([]entity.UserRole, error) {
	roles := []entity.UserRole{}
	err := r.db.SelectContext(ctx, &roles, "SELECT user_id, role_id FROM user_role WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
