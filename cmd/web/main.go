package main

import (
	"log"
	"net/http"

	httpinfra "site-produto/internal/infra/http"
	tpl "site-produto/internal/infra/template"
	"site-produto/internal/page"
	"site-produto/pkg/config"
)

func main() {
	// 1. Carregar configuração
	cfg := config.Load()

	// 2. Inicializar renderer (infra de apresentação)
	renderer, err := tpl.NewRenderer("templates")
	if err != nil {
		log.Fatal(err)
	}

	// 3. Inicializar serviços de domínio
	pageService := page.NewService()

	// 4. Inicializar handlers (HTTP -> domínio -> renderer)
	pageHandler := page.NewHandler(pageService, renderer)

	// 5. Configurar roteamento HTTP explícito
	mux := http.NewServeMux()
	mux.HandleFunc("/", pageHandler.Home)

	// 6. Criar servidor HTTP (infra pura)
	server := httpinfra.NewServer(cfg.Addr, mux)

	// 7. Subir servidor
	log.Println("listening on", cfg.Addr)
	log.Fatal(server.ListenAndServe())
}
