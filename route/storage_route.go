package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterStorageRoutes(r *mux.Router, storageHandler *handler.StorageLocationHandler) {
	r.HandleFunc("/inventory/storageLocation/save", storageHandler.SaveLocation).Methods("POST")
	r.HandleFunc("/inventory/storageLocation/get", storageHandler.GetLocations).Methods("GET")
}
