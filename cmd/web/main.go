package main

import (
	"log"
	"net/http"

	"site-produto/internal/page"
	httpinfra "site-produto/internal/infra/http"
	"site-produto/pkg/config"
)

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/", page.Home)

	server := httpinfra.NewServer(cfg.Addr, mux)

	log.Println("listening on", cfg.Addr)
	log.Fatal(server.ListenAndServe())
}
