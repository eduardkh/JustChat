package handlers

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is an Echo renderer using html/template.
type TemplateRenderer struct {
	templates *template.Template
}

// NewTemplateRenderer parses templates in the templates directory.
func NewTemplateRenderer() *TemplateRenderer {
	t := template.Must(template.ParseGlob("templates/*.html"))
	return &TemplateRenderer{templates: t}
}

// Render renders a template.
func (r *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}
