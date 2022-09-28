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
	user, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Decrypt password
	b64Password, err := encryption.FromBase64(user.Password)
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

	return &model.User{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
