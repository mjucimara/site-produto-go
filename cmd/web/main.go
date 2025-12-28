package main

import (
	"log"
	"net/http"

	httpinfra "site-produto/internal/infra/http"
	tpl "site-produto/internal/infra/template"

	"site-produto/internal/metrics"
	"site-produto/internal/page"

	"site-produto/pkg/config"
)

func main() {

	// ─────────────────────────────────────────────
	// 1. Configuração da aplicação
	// ─────────────────────────────────────────────
	cfg := config.Load()

	// ─────────────────────────────────────────────
	// 2. Infra de apresentação (templates / renderer)
	// ─────────────────────────────────────────────
	renderer, err := tpl.NewRenderer("templates")
	if err != nil {
		log.Fatal(err)
	}

	// ─────────────────────────────────────────────
	// 3. Serviços de domínio (padrão Go: sem construtor vazio)
	// ─────────────────────────────────────────────
	pageService := &page.Service{}
	metricsService := &metrics.Service{}

	// ─────────────────────────────────────────────
	// 4. Handlers (HTTP → domínio → apresentação)
	// ─────────────────────────────────────────────
	pageHandler := page.NewHandler(pageService, renderer)
	metricsHandler := metrics.NewHandler(metricsService, renderer)

	// ─────────────────────────────────────────────
	// 5. Roteamento HTTP explícito
	// ─────────────────────────────────────────────
	mux := http.NewServeMux()
	mux.HandleFunc("/", pageHandler.Home)
	mux.HandleFunc("/metrics", metricsHandler.View)

	// ─────────────────────────────────────────────
	// 6. Servidor HTTP
	// ─────────────────────────────────────────────
	server := httpinfra.NewServer(cfg.Addr, mux)

	// ─────────────────────────────────────────────
	// 7. Inicialização
	// ─────────────────────────────────────────────
	log.Println("listening on", cfg.Addr)
	log.Fatal(server.ListenAndServe())
}
