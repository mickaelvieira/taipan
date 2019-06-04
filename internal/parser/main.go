package parser

import (
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Parse parses the html tree and creates our parsed document
func Parse(URL *url.URL, r io.Reader) (*document.Document, error) {
	document, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var p = Parser{origURL: URL, document: document}
	d := p.Parse()

	return d, nil
}
