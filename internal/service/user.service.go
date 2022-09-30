package service

import (
	"context"

	"github.com/albertopformoso/inventory/encryption"
	"github.com/albertopformoso/inventory/internal/model"
)

func (s *service) RegisterUser(ctx context.Context, email, name, password string) error {
	user, _ := s.repository.GetUserByEmail(ctx, email)
	if user != nil {
		return ErrUserAlreadyExists
	}

	// Hash password
	encryptedPassword, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	encryptedB64Password := encryption.ToBase64(encryptedPassword)
	return s.repository.SaveUser(ctx, email, name, encryptedB64Password)
}

func (s *service) LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	u, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Decrypt password
	b64Password, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}
	decryptedPassword, err := encryption.Decrypt(b64Password)
	if err != nil {
		return nil, err
	}

	if string(decryptedPassword) != password {
		return nil, ErrInvalidCredentials
	}

	user := &model.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}

	return user, nil
}

func (s *service) AddUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repository.GetUserRole(ctx, userID)
	if err != nil {
		return err
	}

	for _, role := range roles {
		if role.RoleID == roleID {
			return ErrUserRoleAlreadyAdded
		}
	}

	return s.repository.SaveUserRole(ctx, userID, roleID)
}

func (s *service) RemoveUserRole(ctx context.Context, userID, roleID int64) error {
	roles, err := s.repository.GetUserRole(ctx, userID)
	if err != nil {
		return err
	}

	var roleFound bool
	for _, role := range roles {
		if role.RoleID == roleID {
			roleFound = true
			break
		}
	}

	if !roleFound {
		return ErrRoleNotFound
	}

	return s.repository.RemoveUserRole(ctx, userID, roleID)
}
