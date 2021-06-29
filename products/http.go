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
	route.HandleFunc("/product/create", Create).Methods("POST")
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
