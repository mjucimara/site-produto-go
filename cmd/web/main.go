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
