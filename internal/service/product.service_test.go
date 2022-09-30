package service

import (
	"testing"

	model "github.com/albertopformoso/inventory/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestAddProduct(t *testing.T) {
	testCases := []struct {
		Name          string
		Product       model.Product
		Email         string
		ExpectedError error
	}{
		{
			Name: "AddProduct_Success",
			Product: model.Product{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       10.00,
			},
			Email:         "admin@email.com",
			ExpectedError: nil,
		},
		{
			Name: "AddProduct_InvalidPermissions",
			Product: model.Product{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       10.00,
			},
			Email:         "customer@email.com",
			ExpectedError: ErrInvalidPermissions,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			// Logic
			err := svc.AddProduct(ctx, tc.Product, tc.Email)

			// Assertions
			assert.Equal(t, tc.ExpectedError, err)
		})
	}
}

func TestGetProducts(t *testing.T) {
	testCases := []struct {
		Name             string
		ID               int64
		ExpectedProducts []model.Product
		ExpectedError    error
	}{
		{
			Name: "GetProducts_Success",
			ExpectedProducts: []model.Product{
				{
					Name:        "Test Product",
					Description: "Test Description",
					Price:       10.00,
				},
			},
			ExpectedError: nil,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			// Logic
			pp, err := svc.GetProducts(ctx)
			var products []model.Product
			for _, p := range pp {
				products = append(products, model.Product{
					Name:        p.Name,
					Description: p.Description,
					Price:       p.Price,
				})
			}

			// Assertions
			assert.Equal(t, tc.ExpectedError, err)
			assert.ElementsMatch(t, tc.ExpectedProducts, products)
		})
	}
}

func TestGetProduct(t *testing.T) {
	testCases := []struct {
		Name            string
		ID              int64
		ExpectedProduct model.Product
		ExpectedError   error
	}{
		{
			Name: "GetProduct_Success",
			ID:   1,
			ExpectedProduct: model.Product{
				Name:        "Test Product",
				Description: "Test Description",
				Price:       10.00,
			},
			ExpectedError: nil,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			// Logic
			p, err := svc.GetProduct(ctx, tc.ID)
			product := model.Product{
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			}

			// Assertions
			assert.Equal(t, tc.ExpectedError, err)
			assert.Equal(t, tc.ExpectedProduct, product)
		})
	}
}
