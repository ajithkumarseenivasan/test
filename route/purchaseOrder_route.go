package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterPurchaseOrderRoutes(r *mux.Router, purchaseOrderHandler *handler.PurchaseOrderHandler) {
	r.HandleFunc("/inventory/purchaseOrder/save", purchaseOrderHandler.SavePurchaseOrder).Methods("POST")
	r.HandleFunc("/inventory/purchaseOrders", purchaseOrderHandler.GetPurchaseOrders).Methods("GET")
}
