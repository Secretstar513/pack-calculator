package main

import (
	"log"
	"net/http"

	"github.com/Secretstar513/pack-calculator/internal/server"
)

func main() {
	packs := server.LoadPacks()
	srv := server.New(packs)
	log.Println("⇢  Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", srv.Routes()))
}
