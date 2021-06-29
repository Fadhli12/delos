package products

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/iannrafisyah/delos/common"
	"github.com/iannrafisyah/delos/products/domain"
	"github.com/iannrafisyah/delos/products/entity"
	"github.com/iannrafisyah/delos/products/handler"
	"github.com/iannrafisyah/delos/products/repository"
)

var handlerProducts domain.ProductsHandler

// Routes :
func Routes(route *mux.Router, db *sql.DB) {
	handlerProducts = handler.NewProductsHandler(
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

	//Read body and Set request value to struct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	product := entity.Products{}
	if err := json.Unmarshal(body, &product); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusInternalServerError),
		})
		return
	}

	if err := handlerProducts.Create(r.Context(), &product); err != nil {
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
	products, err := handlerProducts.List(r.Context())
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

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	productID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	product, err := handlerProducts.Detail(r.Context(), &entity.Products{
		ID: int(productID),
	})
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

	//Read body and Set request value to struct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	product := entity.Products{}
	if err := json.Unmarshal(body, &product); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusInternalServerError),
		})
		return
	}

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	productID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
	}
	product.ID = int(productID)

	if err := handlerProducts.Update(r.Context(), &product); err != nil {
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

	//Get id from segment url
	params := strings.Split(r.URL.Path, "/")
	productID, err := strconv.ParseInt(params[2], 10, 16)
	if err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: common.ErrorRequest(err, http.StatusBadRequest),
		})
		return
	}

	if err := handlerProducts.Delete(r.Context(), &entity.Products{
		ID: int(productID),
	}); err != nil {
		common.ResponseJson(w, &common.ResponseRequest{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.ResponseRequest{})
}
