package page

import tpl "site-produto/internal/infra/template"

type Page struct {
	tpl.PageData
	Description string
}
