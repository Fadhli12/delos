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

// Create :
func Create(w http.ResponseWriter, r *http.Request) {
	if err := handlerProducts.Create(*r); err != nil {
		common.ResponseJson(w, &common.Responses{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.Responses{
		Body: common.ResponseSuccess,
	})
}

// List :
func List(w http.ResponseWriter, r *http.Request) {
	products, err := handlerProducts.List(*r)
	if err != nil {
		common.ResponseJson(w, &common.Responses{
			Error: err,
		})
		return

	}
	common.ResponseJson(w, &common.Responses{
		Body: products,
	})
}

// Detail :
func Detail(w http.ResponseWriter, r *http.Request) {
	product, err := handlerProducts.Detail(*r)
	if err != nil {
		common.ResponseJson(w, &common.Responses{
			Error: err,
		})
		return

	}
	common.ResponseJson(w, &common.Responses{
		Body: product,
	})
}

// Update :
func Update(w http.ResponseWriter, r *http.Request) {
	if err := handlerProducts.Update(*r); err != nil {
		common.ResponseJson(w, &common.Responses{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.Responses{
		Body: common.ResponseSuccess,
	})
}

// Delete :
func Delete(w http.ResponseWriter, r *http.Request) {
	if err := handlerProducts.Delete(*r); err != nil {
		common.ResponseJson(w, &common.Responses{
			Error: err,
		})
		return
	}
	common.ResponseJson(w, &common.Responses{
		Body: common.ResponseSuccess,
	})
}
