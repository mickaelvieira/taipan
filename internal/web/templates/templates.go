package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Templates templates renderer
type Templates struct {
	templates *template.Template
}

// Render implements the echo.Renderer interface
func (r *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

// NewRenderer creates a new renderer
func NewRenderer(dir string) *Templates {
	return &Templates{
		templates: template.Must(template.New("html-tmpl").ParseGlob(dir + "/*.html")),
	}
}
