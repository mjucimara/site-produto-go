package template

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
)

type Renderer struct {
	base *template.Template
	dir  string
}

func NewRenderer(dir string) (*Renderer, error) {
	base, err := template.ParseFiles(filepath.Join(dir, "layout.html"))
	if err != nil {
		return nil, err
	}

	return &Renderer{
		base: base,
		dir:  dir,
	}, nil
}

func (r *Renderer) Render(w http.ResponseWriter, name string, data any) {
	tpl, err := r.base.Clone()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tpl.ParseFiles(filepath.Join(r.dir, name+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write(buf.Bytes())
}
