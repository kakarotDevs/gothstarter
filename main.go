package main

import (
	"fmt"
	"gothstarter/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	// Middleware should come before route declarations
	router.Use(middleware.Recoverer)

	router.Handle("/public/*", public())
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Handle("/login", handlers.Make(handlers.HandleLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":3000"
	}

	slog.Info("3, 2, 1, Let's Jam...", "url", fmt.Sprintf("http://localhost%s", listenAddr))
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
