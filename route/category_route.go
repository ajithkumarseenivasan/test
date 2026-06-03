package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(r *mux.Router, categoryHandler *handler.CategoryHandler) {
	r.HandleFunc("/category/save", categoryHandler.SaveCategory).Methods("POST")
}
