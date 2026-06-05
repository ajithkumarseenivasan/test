package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(r *mux.Router, categoryHandler *handler.CategoryHandler) {
	r.HandleFunc("/inventory/category/save", categoryHandler.SaveCategory).Methods("POST")
	r.HandleFunc("/inventory/getCategory", categoryHandler.GetCategories).Methods("GET")
}
