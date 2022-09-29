package service

import (
	"os"
	"testing"

	"github.com/albertopformoso/inventory/encryption"
	"github.com/albertopformoso/inventory/internal/entity"
	"github.com/albertopformoso/inventory/internal/repository"
	"github.com/stretchr/testify/mock"
)

var repo *repository.MockUser
var svc User

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	user := &entity.User{
		Email:    "test@exists.com",
		Password: encryptedPassword,
	}

	repo = &repository.MockUser{}
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(user, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("GetUserRole", mock.Anything, int64(1)).Return([]entity.UserRole{{UserID: 1, RoleID: 1}}, nil)
	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	svc = New(repo)

	code := m.Run()
	os.Exit(code)
}
