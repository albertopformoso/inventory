package service

import (
	"context"

	"github.com/albertopformoso/inventory/internal/model"
)

var (
	validRolesToAddProduct []int64 = []int64{1, 2}
)

func (s *service) AddProduct(ctx context.Context, product model.Product, email string) error {
	u, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	rr, err := s.repository.GetUserRole(ctx, u.ID)
	if err != nil {
		return err
	}

	var userCanAdd bool
	for _, r := range rr {
		for _, vr := range validRolesToAddProduct {
			if vr == r.RoleID {
				userCanAdd = true
			}
		}
	}

	if !userCanAdd {
		return ErrInvalidPermissions
	}

	return s.repository.SaveProduct(
		ctx,
		product.Name,
		product.Description,
		product.Price,
		u.ID,
	)
}

func (s *service) GetProducts(ctx context.Context) ([]model.Product, error) {
	pp, err := s.repository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := []model.Product{}
	for _, p := range pp {
		products = append(products, model.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return products, nil
}
func (s *service) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	p, err := s.repository.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	product := &model.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}

	return product, nil
}
