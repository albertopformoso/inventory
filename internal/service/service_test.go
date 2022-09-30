package service

import (
	"os"
	"testing"

	"github.com/albertopformoso/inventory/encryption"
	"github.com/albertopformoso/inventory/internal/entity"
	"github.com/albertopformoso/inventory/internal/repository"
	"github.com/stretchr/testify/mock"
)

var repo *repository.MockRepository
var svc Service

func TestMain(m *testing.M) {
	validPassword, _ := encryption.Encrypt([]byte("validPassword"))
	encryptedPassword := encryption.ToBase64(validPassword)
	user := &entity.User{
		Email:    "test@exists.com",
		Password: encryptedPassword,
	}
	adminUser := &entity.User{
		ID:       1,
		Email:    "admin@email.com",
		Password: encryptedPassword,
	}
	customerUser := &entity.User{
		ID:       2,
		Email:    "customer@email.com",
		Password: encryptedPassword,
	}
	testProducts := []entity.Product{
		{
			ID:          1,
			Name:        "Test Product",
			Description: "Test Description",
			Price:       10.00,
		},
	}

	repo = &repository.MockRepository{}
	// User
	repo.On("GetUserByEmail", mock.Anything, "test@test.com").Return(nil, nil)
	repo.On("GetUserByEmail", mock.Anything, "test@exists.com").Return(user, nil)
	repo.On("GetUserByEmail", mock.Anything, "admin@email.com").Return(adminUser, nil)
	repo.On("GetUserByEmail", mock.Anything, "customer@email.com").Return(customerUser, nil)
	repo.On("SaveUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	repo.On("GetUserRole", mock.Anything, int64(1)).Return([]entity.UserRole{{UserID: 1, RoleID: 1}}, nil)
	repo.On("GetUserRole", mock.Anything, int64(2)).Return([]entity.UserRole{{UserID: 2, RoleID: 3}}, nil)
	repo.On("SaveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("RemoveUserRole", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Product
	repo.On("SaveProduct", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	repo.On("GetProducts", mock.Anything).Return(testProducts, nil)
	repo.On("GetProduct", mock.Anything, int64(1)).Return(&entity.Product{ID: 1, Name: "Test Product", Description: "Test Description", Price: 10.00}, nil)

	svc = New(repo)

	code := m.Run()
	os.Exit(code)
}
