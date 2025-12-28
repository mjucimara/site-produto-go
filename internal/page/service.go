package page

import tpl "site-produto/internal/infra/template"

type Service struct{}

func (s *Service) Home() Page {
	return Page{
		PageData: tpl.PageData{
			Title: "Site Produto",
		},
		Description: "Este é exatamente o site que você vai receber.",
	}
}
