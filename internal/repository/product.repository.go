package repository

import (
	context "context"
	"errors"

	entity "github.com/albertopformoso/inventory/internal/entity"
	"github.com/albertopformoso/inventory/internal/helper"
)

const (
	queryInsertProduct = `
	INSERT INTO product (name, description, price, created_by) VALUES (?, ?, ?, ?);`
	queryGetAllProducts = `
	SELECT id, name, description, price, created_by FROM product;`
	queryGetProductByID = `
	SELECT id, name, description, price, created_by
	FROM product
	WHERE id = ?`
)

func (r *repository) SaveProduct(ctx context.Context, name, description string, price float32, createdBy int64) (err error) {
	defer helper.PanicRecover(func(e interface{}) {
		r.log.Error().Msgf("panicRecovered - error:", e)

		err = errors.New("unexpected Failed to Save Product - recovered into unhandled error")
	})

	_, err = r.db.ExecContext(ctx, queryInsertProduct, name, description, price, createdBy)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetProducts(ctx context.Context) ([]entity.Product, error) {
	pp := []entity.Product{}
	err := r.db.SelectContext(ctx, &pp, queryGetAllProducts)
	if err != nil {
		return nil, err
	}

	return pp, nil
}
func (r *repository) GetProduct(ctx context.Context, id int64) (*entity.Product, error) {
	pp := []entity.Product{}
	err := r.db.SelectContext(ctx, &pp, queryGetProductByID, id)
	if err != nil {
		return nil, err
	}

	p := &pp[0]
	return p, nil
}
