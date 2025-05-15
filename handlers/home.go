package handlers

import (
	"gothstarter/views/home"
	"log/slog"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	slog.Info("Serving homepage")
	return Render(w, r, home.Index())
}
