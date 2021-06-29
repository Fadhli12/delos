package domain

import (
	"context"
	"net/http"

	"github.com/iannrafisyah/delos/products/entity"
)

// ProductsHandler :
type ProductsHandler interface {
	Create(http.Request) error
}

// ProductsRepository :
type ProductsRepository interface {
	Create(context.Context, *entity.Products) error
}
