package templates

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// Template templates renderer
type Template struct {
	templates *template.Template
}

// Render implements the echo.Renderer interface
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewRenderer creates a new renderer
func NewRenderer(dir string) *Template {
	return &Template{
		templates: template.Must(template.New("html-tmpl").ParseGlob(dir + "/*.html")),
	}
}
