package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, userHandler *handler.UserHandler) {
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{name}", userHandler.GetUserByName).Methods("GET")
	// r.HandleFunc("/users/save", userHandler.SaveNewUser).Methods("POST")
}
