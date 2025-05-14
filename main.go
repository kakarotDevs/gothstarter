package main

import (
	"gothstarter/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Handle("/public/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("3, 2, 1, Let's Jam...", "url", "http://localhost"+listenAddr)
	http.ListenAndServe(listenAddr, router)
}
