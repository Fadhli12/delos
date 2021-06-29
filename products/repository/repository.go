package repository

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/iannrafisyah/delos/common"
	"github.com/iannrafisyah/delos/products/domain"
	"github.com/iannrafisyah/delos/products/entity"
)

type ProductsRepository struct {
	conn *sql.DB
}

const (
	Create = `INSERT INTO "products" (title,description,image,rating) VALUES ($1,$2,$3,$4)`
)

// NewProductsRepository :
func NewProductsRepository(db *sql.DB) domain.ProductsRepository {
	return &ProductsRepository{
		conn: db,
	}
}

// Create
func (r *ProductsRepository) Create(ctx context.Context, product *entity.Products) error {
	if err := r.conn.QueryRowContext(ctx, Create, product.Title, product.Description, product.Image, product.Rating).Err(); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}
	return nil
}
