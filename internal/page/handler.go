package page

import "net/http"

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	page := h.service.Home()
	w.Write([]byte(page.Title + "\n" + page.Description))
}
