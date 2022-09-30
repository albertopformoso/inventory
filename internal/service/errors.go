package service

import "errors"

var (
	ErrUserAlreadyExists    = errors.New("user already exists")
	ErrInvalidCredentials   = errors.New("invalid password")
	ErrUserRoleAlreadyAdded = errors.New("role already added for this user")
	ErrRoleNotFound         = errors.New("role not found")
	ErrInvalidPermissions   = errors.New("user does not have permission to add product")
)
