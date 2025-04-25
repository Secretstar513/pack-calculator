package main

import (
	"net/http"

	"github.com/Secretstar513/pack-calculator/internal/server"
)

var srv http.Handler

func init() {
	packs := server.LoadPacks()
	srv = server.New(packs).Routes() // Build routes just once when cold-started
}

// Vercel entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	srv.ServeHTTP(w, r)
}
