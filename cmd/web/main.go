package main

import (
	"log"
	"net/http"

	httpinfra "site-produto/internal/infra/http"
	tpl "site-produto/internal/infra/template"

	"site-produto/internal/health"
	"site-produto/internal/metrics"
	"site-produto/internal/page"

	"site-produto/pkg/config"
)

func main() {

	// ─────────────────────────────────────────────
	// 1. Configuração
	// ─────────────────────────────────────────────
	cfg := config.Load()

	// ─────────────────────────────────────────────
	// 2. Renderer (templates)
	// ─────────────────────────────────────────────
	renderer, err := tpl.NewRenderer("templates")
	if err != nil {
		log.Fatal(err)
	}

	// ─────────────────────────────────────────────
	// 3. Serviços de domínio
	// ─────────────────────────────────────────────
	pageService := &page.Service{}
	metricsService := &metrics.Service{}

	// ─────────────────────────────────────────────
	// 4. Handlers
	// ─────────────────────────────────────────────
	pageHandler := page.NewHandler(pageService, renderer)
	metricsHandler := metrics.NewHandler(metricsService, renderer)
	healthHandler := health.NewHandler()

	// ─────────────────────────────────────────────
	// 5. Roteamento HTTP
	// ─────────────────────────────────────────────
	mux := http.NewServeMux()

	mux.HandleFunc("/", pageHandler.Home)
	mux.HandleFunc("/metrics", metricsHandler.View)
	mux.HandleFunc("/health", healthHandler.Check)

	// arquivos estáticos
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

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
