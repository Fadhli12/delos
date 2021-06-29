package domain

import (
	"context"
	"net/http"

	"github.com/iannrafisyah/delos/products/entity"
)

// ProductsHandler :
type ProductsHandler interface {
	Create(http.Request) error
	List(http.Request) ([]*entity.Products, error)
	Detail(http.Request) (*entity.Products, error)
	Update(http.Request) error
	Delete(http.Request) error
}

// ProductsRepository :
type ProductsRepository interface {
	Create(context.Context, *entity.Products) error
	List(context.Context) ([]*entity.Products, error)
	Detail(context.Context, *entity.Products) (*entity.Products, error)
	Update(context.Context, *entity.Products) error
	Delete(context.Context, *entity.Products) error
}
