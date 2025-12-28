package page

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

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	page := h.service.Home()
	h.renderer.Render(w, "home", page)
}
