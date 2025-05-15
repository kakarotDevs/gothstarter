//go:build dev

package main

import (
	"fmt"
	"net/http"
	"os"
)

func public() http.Handler {
	fmt.Println("Using local public/ directory for static files (dev mode)")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}
