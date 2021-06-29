package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/iannrafisyah/delos/common"
	"github.com/iannrafisyah/delos/products/domain"
	"github.com/iannrafisyah/delos/products/entity"
)

type ProductsHandler struct {
	Repository domain.ProductsRepository
}

// NewProductsHandler :
func NewProductsHandler(productRepository domain.ProductsRepository) domain.ProductsHandler {
	return &ProductsHandler{
		Repository: productRepository,
	}
}

func (h *ProductsHandler) Create(request http.Request) error {

	rating, err := strconv.ParseInt(request.FormValue("rating"), 10, 16)
	if err != nil {
		return common.ErrorRequest(err, http.StatusBadRequest)
	}

	product := &entity.Products{
		Title:       request.FormValue("title"),
		Description: request.FormValue("description"),
		Rating:      int(rating),
		Image:       request.FormValue("image"),
	}

	//Validation
	if product.Title == "" {
		return common.ErrorRequest(errors.New(common.TitleRequired), http.StatusBadRequest)
	} else if product.Description == "" {
		return common.ErrorRequest(errors.New(common.DescriptionRequired), http.StatusBadRequest)
	} else if product.Rating < 1 || product.Rating > 5 {
		return common.ErrorRequest(errors.New(common.RatingRequired), http.StatusBadRequest)
	} else if product.Image == "" {
		return common.ErrorRequest(errors.New(common.ImageRequired), http.StatusBadRequest)
	}

	//Create product
	if err := h.Repository.Create(request.Context(), product); err != nil {
		return err
	}

	return nil
}
