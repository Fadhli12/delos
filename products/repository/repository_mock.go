package repository

import (
	"context"
	"errors"

	"github.com/iannrafisyah/delos/products/entity"
	"github.com/stretchr/testify/mock"
)

type ProductsRepositoryMock struct {
	Mock mock.Mock
}

// Create :
func (r *ProductsRepositoryMock) Create(ctx context.Context, products *entity.Products) error {
	args := r.Mock.Called(products)
	if args.Get(0) == nil {
		return errors.New("Failed create a product")
	}
	return nil
}

// List :
func (r *ProductsRepositoryMock) List(ctx context.Context) ([]*entity.Products, error) {
	return nil, nil
}

// Detail :
func (r *ProductsRepositoryMock) Detail(ctx context.Context, products *entity.Products) (*entity.Products, error) {
	args := r.Mock.Called(products)
	if args.Get(0) == nil {
		return nil, errors.New("Failed create a product")
	}

	fetchProduct := args.Get(0).(*entity.Products)
	if fetchProduct.ID != products.ID {
		return nil, errors.New("Product not found")
	}
	return fetchProduct, nil
}

// Update :
func (r *ProductsRepositoryMock) Update(ctx context.Context, products *entity.Products) error {
	_, err := r.Detail(ctx, products)
	if err != nil {
		return err
	}

	args := r.Mock.Called(products)
	if args.Get(0) == nil {
		return errors.New("Failed update a product")
	}
	return nil
}

// Delete :
func (r *ProductsRepositoryMock) Delete(ctx context.Context, products *entity.Products) error {
	_, err := r.Detail(ctx, products)
	if err != nil {
		return err
	}

	args := r.Mock.Called(products)
	if args.Get(0) == nil {
		return errors.New("Failed delete a product")
	}

	return nil
}
