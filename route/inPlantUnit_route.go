package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterInPlantUnitRoutes(r *mux.Router, inPlantUnitHandler *handler.InPlantUnitHandler) {
	r.HandleFunc("/inventory/inPlantUnit/save", inPlantUnitHandler.SaveInPlantUnit).Methods("POST")
	r.HandleFunc("/inventory/getInPlantUnits", inPlantUnitHandler.GetInPlantUnits).Methods("GET")
}
