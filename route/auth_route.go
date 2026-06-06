package route

import (
	"user-management/handler"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router, authHandler *handler.AuthHandler) {
	r.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
}

func RegisterProtectedAuthRoutes(r *mux.Router, authHandler *handler.AuthHandler) {
	r.HandleFunc("/auth/me", authHandler.Me).Methods("GET")
}
