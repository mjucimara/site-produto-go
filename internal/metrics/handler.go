package metrics

import (
	"net/http"

	tpl "site-produto/internal/infra/template"
)

type Handler struct {
	service  *Service
	renderer *tpl.Renderer
}

func NewHandler(s *Service, r *tpl.Renderer) *Handler {
	return &Handler{
		service:  s,
		renderer: r,
	}
}

func (h *Handler) View(w http.ResponseWriter, r *http.Request) {
	data := struct {
		tpl.PageData
		Metrics []Metric
	}{
		PageData: tpl.PageData{
			Title: "Métricas",
		},
		Metrics: h.service.All(),
	}

	h.renderer.Render(w, "metrics", data)
}
