package parser

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/url"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// Parse parses the html tree and creates our parsed document
func Parse(URL *url.URL, r io.Reader, findFeeds bool) (*document.Document, error) {
	document, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Parsing RSS feeds too? [%t]\n", findFeeds)

	var p = Parser{origURL: URL, document: document}
	if findFeeds {
		p.ShouldFindSyndicationSource()
	}
	return p.Parse(), nil
}
