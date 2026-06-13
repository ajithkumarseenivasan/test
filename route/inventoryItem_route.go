package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterInventoryItemRoutes(r *mux.Router, inventoryItemHandler *handler.InventoryItemHandler) {
	r.HandleFunc("/inventory/invItem/save", inventoryItemHandler.SaveInventoryItem).Methods("POST")
	r.HandleFunc("/inventory/invItem/get", inventoryItemHandler.GetInventoryItems).Methods("GET")
}
