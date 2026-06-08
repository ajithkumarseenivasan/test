package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterVendorRoutes(r *mux.Router, vendorHandler *handler.VendorHandler) {
	r.HandleFunc("/inventory/vendor/save", vendorHandler.SaveVendor).Methods("POST")
	r.HandleFunc("/inventory/vendor/get", vendorHandler.GetVendors).Methods("GET")
}
