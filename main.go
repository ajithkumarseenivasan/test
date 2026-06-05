package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-management/app"
	"user-management/route"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	application := app.NewApplication()

	r := mux.NewRouter()

	route.RegisterUserRoutes(r, application.UserHandler)
	route.RegisterCategoryRoutes(r, application.CategoryHandler)

	ser := &http.Server{
		Addr:    ":8080",
		Handler: corsMiddleware(r),
	}

	go func() {
		log.Println("Server running on :8080")
		if err := ser.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down the server...!")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := ser.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exiting")
}
