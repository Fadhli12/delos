package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

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

// Create :
func (h *ProductsHandler) Create(request http.Request) error {

	//Read body and Set request value to struct
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return common.ErrorRequest(err, http.StatusBadRequest)
	}

	product := entity.Products{}
	if err := json.Unmarshal(body, &product); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
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
	if err := h.Repository.Create(request.Context(), &product); err != nil {
		return err
	}

	return nil
}

// List :
func (h *ProductsHandler) List(request http.Request) ([]*entity.Products, error) {

	//Get list products
	products, err := h.Repository.List(request.Context())
	if err != nil {
		return nil, err
	}
	return products, nil
}

// Detail :
func (h *ProductsHandler) Detail(request http.Request) (*entity.Products, error) {

	//Get id from segment url
	params := strings.Split(request.URL.Path, "/")
	productID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		return nil, common.ErrorRequest(err, http.StatusBadRequest)
	}

	//Get detail product
	product, err := h.Repository.Detail(request.Context(), &entity.Products{
		ID: int(productID),
	})
	if err != nil {
		return nil, err
	}
	return product, nil
}

// Update :
func (h *ProductsHandler) Update(request http.Request) error {

	//Read body and Set request value to struct
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return common.ErrorRequest(err, http.StatusBadRequest)
	}

	product := entity.Products{}
	if err := json.Unmarshal(body, &product); err != nil {
		return common.ErrorRequest(err, http.StatusInternalServerError)
	}

	//Get id from segment url
	params := strings.Split(request.URL.Path, "/")
	productID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		return common.ErrorRequest(err, http.StatusBadRequest)
	}
	product.ID = int(productID)

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

	//Update product
	if err := h.Repository.Update(request.Context(), &product); err != nil {
		return err
	}
	return nil
}

// Delete :
func (h *ProductsHandler) Delete(request http.Request) error {

	//Get id from segment url
	params := strings.Split(request.URL.Path, "/")
	productID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		return common.ErrorRequest(err, http.StatusBadRequest)
	}

	//Delete product
	if err := h.Repository.Delete(request.Context(), &entity.Products{
		ID: int(productID),
	}); err != nil {
		return err
	}
	return nil
}
