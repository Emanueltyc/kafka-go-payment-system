package email

import (
	"bytes"
	"embed"
	"html/template"
)

type Renderer struct {
	templates *template.Template
}

//go:embed templates/*.html
var templateFS embed.FS

func NewRenderer() (*Renderer, error) {
	tmpl, err := template.ParseFS(
		templateFS,
		"templates/base.html",
		"templates/*.html",
	)
	if err != nil {
		return nil, err
	}

	return &Renderer{templates: tmpl}, nil
}

func (r *Renderer) Render(templateName string, data any) (string, error) {
	var buf bytes.Buffer

	err := r.templates.ExecuteTemplate(&buf, templateName, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
