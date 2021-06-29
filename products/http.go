package products

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iannrafisyah/delos/common"
	"github.com/iannrafisyah/delos/products/domain"
	"github.com/iannrafisyah/delos/products/handlers"
	"github.com/iannrafisyah/delos/products/repository"
)

var handlerProducts domain.ProductsHandler

// Routes :
func Routes(route *mux.Router, db *sql.DB) {
	handlerProducts = handlers.NewProductsHandler(
		repository.NewProductsRepository(db),
	)
	route.HandleFunc("/products", Create).Methods("POST")
	route.HandleFunc("/products", List).Methods("GET")
	route.HandleFunc("/products/{id:[0-9]+}", Detail).Methods("GET")
	route.HandleFunc("/products/{id:[0-9]+}", Update).Methods("PUT")
	route.HandleFunc("/products/{id:[0-9]+}", Delete).Methods("DELETE")
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags product
// @Accept json
// @Produce json
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param rating body int true "Rating"
// @Param image body string true "Image"
// @Success 200 {object} common.Responses
// @Failure 400,500 {object} common.Responses
// @Router /products [post]
func Create(w http.ResponseWriter, r *http.Request) {
	if err := handlerProducts.Create(*r); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}

// ListProduct godoc
// @Summary List products
// @Description Returns product data
// @Tags product
// @Accept json
// @Produce json
// @Success 200 {object} common.Responses
// @Failure 500 {object} common.Responses
// @Router /products [get]
func List(w http.ResponseWriter, r *http.Request) {
	products, err := handlerProducts.List(*r)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return

	}
	common.ResponseJson(w, &common.ResponseRequest{
		Data: products,
	})
}

// DetailProduct godoc
// @Summary Detail a product
// @Description Return detail product with param id
// @Tags product
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} common.Responses
// @Failure 400,404,500 {object} common.Responses
// @Router /products/{id} [get]
func Detail(w http.ResponseWriter, r *http.Request) {
	product, err := handlerProducts.Detail(*r)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return

	}
	common.ResponseJson(w, &common.ResponseRequest{
		Data: product,
	})
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product with the input payload and param id
// @Tags product
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Param title body string true "Title"
// @Param description body string true "Description"
// @Param rating body int true "Rating"
// @Param image body string true "Image"
// @Success 200 {object} common.Responses
// @Failure 400,500 {object} common.Responses
// @Router /products/{id} [put]
func Update(w http.ResponseWriter, r *http.Request) {
	if err := handlerProducts.Update(*r); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product with param id
// @Tags product
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} common.Responses
// @Failure 400,404,500 {object} common.Responses
// @Router /products/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	if err := handlerProducts.Delete(*r); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}
