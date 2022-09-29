package service

import (
	context "context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var ctx context.Context = context.Background()

func TestRegisterUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		UserName      string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "RegisterUser_Success",
			Email:         "test@test.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "RegisterUser_UserAlreadyExists",
			Email:         "test@exists.com",
			UserName:      "test",
			Password:      "validPassword",
			ExpectedError: ErrUserAlreadyExists,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			// Logic
			err := svc.RegisterUser(ctx, tc.Email, tc.Name, tc.Password)

			// Assertions
			assert.Equal(t, tc.ExpectedError, err)
		})
	}
}

func TestLoginUser(t *testing.T) {
	testCases := []struct {
		Name          string
		Email         string
		Password      string
		ExpectedError error
	}{
		{
			Name:          "LoginUser_Success",
			Email:         "test@exists.com",
			Password:      "validPassword",
			ExpectedError: nil,
		},
		{
			Name:          "LoginUser_InvalidPassword",
			Email:         "test@exists.com",
			Password:      "invalidPassword",
			ExpectedError: ErrInvalidCredentials,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			// Logic
			_, err := svc.LoginUser(ctx, tc.Email, tc.Password)

			// Assertions
			assert.Equal(t, tc.ExpectedError, err)
		})
	}
}

func TestAddUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			Name:          "AddUserRole_Success",
			UserID:        1,
			RoleID:        2,
			ExpectedError: nil,
		},
		{
			Name:          "AddUserRole_UserAlreadyHasRole",
			UserID:        1,
			RoleID:        1,
			ExpectedError: ErrUserRoleAlreadyAdded,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			// Logic
			err := svc.AddUserRole(ctx, tc.UserID, tc.RoleID)

			// Assertions
			assert.Equal(t, tc.ExpectedError, err)
		})
	}
}

func TestRemoveUserRole(t *testing.T) {
	testCases := []struct {
		Name          string
		UserID        int64
		RoleID        int64
		ExpectedError error
	}{
		{
			Name:          "RemoveUserRole_Success",
			UserID:        1,
			RoleID:        1,
			ExpectedError: nil,
		},
		{
			Name:          "RemoveUserRole_UserDoesNotHaveRole",
			UserID:        1,
			RoleID:        3,
			ExpectedError: ErrRoleNotFound,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			// Logic
			err := svc.RemoveUserRole(ctx, tc.UserID, tc.RoleID)

			// Assertions
			assert.Equal(t, tc.ExpectedError, err)
		})
	}
}
