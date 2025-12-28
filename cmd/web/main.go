package main

import (
	"log"
	"net/http"

	httpinfra "site-produto/internal/infra/http"
	"site-produto/internal/page"
	"site-produto/pkg/config"
)

func main() {
	// 1. Carregar configuração
	cfg := config.Load()

	// 2. Inicializar serviços de domínio
	pageService := page.NewService()

	// 3. Inicializar handlers (tradução HTTP -> domínio)
	pageHandler := page.NewHandler(pageService)

	// 4. Configurar roteamento HTTP explícito
	mux := http.NewServeMux()
	mux.HandleFunc("/", pageHandler.Home)

	// 5. Criar servidor HTTP (infra pura)
	server := httpinfra.NewServer(cfg.Addr, mux)

	// 6. Subir servidor
	log.Println("listening on", cfg.Addr)
	log.Fatal(server.ListenAndServe())
}
