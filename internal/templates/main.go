package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Renderer templates renderer
type Renderer struct {
	templates *template.Template
}

// Render implements the echo.Renderer interface
func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

// NewRenderer creates a new renderer
func NewRenderer(dir string) *Renderer {
	return &Renderer{
		templates: template.Must(template.New("html-tmpl").ParseGlob(dir + "/*.html")),
	}
}
