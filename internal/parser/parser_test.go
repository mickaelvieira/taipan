package parser

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"

	"github.com/PuerkitoBio/goquery"
)

var origURL *url.URL

func init() {
	var err error
	origURL, err = url.FromRawURL("https://foo.bar")
	if err != nil {
		log.Fatal(err)
	}
}

func getDocumentWithHead(head string) *goquery.Document {
	html := fmt.Sprintf("<html><head>%s</head><body></body>", head)
	return getDocument(html)
}

func getDocument(html string) *goquery.Document {
	document, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Fatal(err)
	}
	return document
}

func TestGetWordpressFeed(t *testing.T) {
	var p = Parser{origURL: origURL}
	var s = p.getWordpressFeed()
	var e = origURL.String() + "/feed/"

	if s[0].URL.String() != e {
		t.Errorf("Incorrect URL: Wanted %s; got %s", e, s[0].URL.String())
	}
	if s[0].Title != defaultWPFeedTitle {
		t.Errorf("Incorrect Title: Wanted %s; got %s", defaultWPFeedTitle, s[0].Title)
	}
	if s[0].Type != syndication.RSS {
		t.Errorf("Incorrect Type: Wanted %s; got %s", syndication.RSS, s[0].Title)
	}
}

func TestDocumentURL(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"", origURL.String()}, // origin URL
		{"<link rel=\"canonical\" href=\"https://bar.foo\"/>", "https://bar.foo"},                       // canonical URL
		{"<meta property=\"og:url\" content=\"https://bar.foo/facebook\">", "https://bar.foo/facebook"}, // facebook URL
		{"<meta name=\"twitter:url\" content=\"https://bar.foo/twitter\">", "https://bar.foo/twitter"},  // twitter URL
		{"<link rel=\"canonical\" href=\"https://bar.foo\"/>" +
			"<meta name=\"twitter:url\" content=\"https://bar.foo/facebook\">" +
			"<meta name=\"twitter:url\" content=\"https://bar.foo/twitter\">", "https://bar.foo"}, // canonical URL precedence
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Document URL [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var p = Parser{origURL: origURL, document: getDocumentWithHead(tc.i)}
			var d = p.Parse()
			var e = tc.o
			if d.URL.String() != e {
				t.Errorf("Incorrect document URL: Wanted %s; got %s", e, d.Title)
			}
		})
	}
}

func TestDocumentTitle(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"", ""},                            // no title
		{"<title>   Foo   </title>", "Foo"}, // html title
		{"<meta property=\"og:title\" content=\" Foo Bar \">", "Foo Bar"}, // facebook title
		{"<meta name=\"twitter:title\" content=\"baz\">", "baz"},          // twitter title
		{"<title>Foo</title>" +
			"<meta name=\"og:title\" content=\"baz\">" +
			"<meta name=\"twitter:title\" content=\"baz\">", "Foo"}, // html title precedence
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Document title [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var p = Parser{origURL: origURL, document: getDocumentWithHead(tc.i)}
			var d = p.Parse()
			var e = tc.o
			if d.Title != e {
				t.Errorf("Incorrect document Title: Wanted %s; got %s", e, d.Title)
			}
		})
	}
}

func TestDocumentDescription(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"", ""}, // no description
		{"<meta name=\"description\" content=\" Foo  \">", "Foo"},               // html description
		{"<meta property=\"og:description\" content=\" Foo Bar \">", "Foo Bar"}, // facebook description
		{"<meta name=\"twitter:description\" content=\"baz\">", "baz"},          // twitter description
		{"<meta name=\"description\" content=\"Foo\">" +
			"<meta name=\"og:description\" content=\"baz\">" +
			"<meta name=\"twitter:description\" content=\"baz\">", "Foo"}, // html description precedence
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Document description [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var p = Parser{origURL: origURL, document: getDocumentWithHead(tc.i)}
			var d = p.Parse()
			var e = tc.o
			if d.Description != e {
				t.Errorf("Incorrect document Description: Wanted %s; got %s", e, d.Description)
			}
		})
	}
}

func TestDocumentCharset(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"", ""},                                 // no charset
		{"<meta charset=\" utf-8  \">", "utf-8"}, // html5 meta charset
		{"<meta http-equiv=\"Content-Type\" content=\"text/html; charset=iso-8859-1\" />", "iso-8859-1"}, // html4 meta charset
		{"<meta charset=\"utf-8\"> " +
			"<meta http-equiv=\"Content-Type\" content=\"text/html; charset=iso-8859-1\" />", "utf-8"}, // html5 charset precedence
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Document charset [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var p = Parser{origURL: origURL, document: getDocumentWithHead(tc.i)}
			var d = p.Parse()
			var e = tc.o
			if d.Charset != e {
				t.Errorf("Incorrect document Charset: Wanted %s; got %s", e, d.Charset)
			}
		})
	}
}

func TestDocumentLang(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"<html><head></head>", ""},                                                         // no lang
		{"<html lang=\"fr\"><head></head>", "fr"},                                           // html5 lang
		{"<html><head><meta http-equiv=\"content-language\" content=\"de\"/></head>", "de"}, // html4 lang
		{"<html lang=\"fr\">" +
			"<head><meta http-equiv=\"content-language\" content=\"de\"/></head></head>", "fr"}, // html5 lang precedence
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Document charset [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var p = Parser{origURL: origURL, document: getDocument(tc.i)}
			var d = p.Parse()
			var e = tc.o
			if d.Lang != e {
				t.Errorf("Incorrect document Lang: Wanted %s; got %s", e, d.Lang)
			}
		})
	}
}

func TestDocumentImage(t *testing.T) {
	var testcase = []struct {
		i string
		o string
	}{
		{"", ""}, // no image
		{"<meta property=\"og:image\" content=\"https://bar.foo/facebook.jpg\">", "https://bar.foo/facebook.jpg"},    // facebook description
		{"<meta name=\"twitter:image\" content=\"https://bar.foo/twitter.jpg\">", "https://bar.foo/twitter.jpg"},     // twitter description
		{"<meta name=\"twitter:image:src\" content=\"https://bar.foo/twitter.jpg\">", "https://bar.foo/twitter.jpg"}, // twitter description
		{"<meta property=\"og:image\" content=\"https://bar.foo/facebook.jpg\">" +
			"<meta name=\"twitter:image\" content=\"https://bar.foo/twitter.jpg\">", "https://bar.foo/facebook.jpg"}, // html description precedence
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Document description [%d]", idx)
		t.Run(name, func(t *testing.T) {
			var p = Parser{origURL: origURL, document: getDocumentWithHead(tc.i)}
			var d = p.Parse()
			var e = tc.o
			if e == "" {
				if d.Image != nil {
					t.Errorf("Incorrect document Image: Wanted nil; got %s", d.Image.URL.String())
				}
			} else {
				if d.Image.URL.String() != e {
					t.Errorf("Incorrect document Image: Wanted %s; got %s", e, d.Image.URL.String())
				}
			}
		})
	}
}

func TestDocumentFeeds(t *testing.T) {
	var html = "<html><head>" +
		"<link href=\"https://bar.foo/rss\" type=\"application/rss+xml\" title=\"RSS feed\" >" +
		"<link href=\"https://bar.foo/atom\" >" +
		"<link href=\"https://bar.foo/atom\" type=\"application/atom+xml\" title=\"Atom feed\" >" +
		"<link type=\"application/rss+xml\" >" +
		"<link href=\"https://bar.foo/rss\" type=\"application/rss+xml\" >" +
		"<link href=\"https://bar.foo/atom\" type=\"application/atom+xml\" >" +
		"</head>" +
		"<body></body>"
	var p = Parser{origURL: origURL, document: getDocument(html)}
	p.ShouldFindSyndicationSource()
	var d = p.Parse()

	var testcase = []struct {
		h  string
		tp string
		t  string
	}{
		{"https://bar.foo/rss", string(syndication.RSS), "RSS feed"},
		{"https://bar.foo/atom", string(syndication.ATOM), "Atom feed"},
		{"https://bar.foo/rss", string(syndication.RSS), ""},
		{"https://bar.foo/atom", string(syndication.ATOM), ""},
	}

	if len(d.Feeds) != 4 {
		t.Errorf("Incorrect number of feeds: Wanted %d; got %d", 4, len(d.Feeds))
	}

	for idx, tc := range testcase {
		name := fmt.Sprintf("Document feeds [%d]", idx)
		f := d.Feeds[idx]
		t.Run(name, func(t *testing.T) {
			if f.URL.String() != tc.h {
				t.Errorf("Incorrect feed URL: Wanted %s; got %s", tc.h, f.URL.String())
			}
			if string(f.Type) != tc.tp {
				t.Errorf("Incorrect document Lang: Wanted %s; got %s", tc.tp, f.Type)
			}
			if f.Title != tc.t {
				t.Errorf("Incorrect document Lang: Wanted %s; got %s", tc.t, f.Title)
			}
		})
	}
}

func TestDocumentWPFeeds(t *testing.T) {
	var html = "<html><head>" +
		"<link href=\"https://bar.foo/rss\" type=\"application/rss+xml\" title=\"RSS feed \" >" +
		"<link href=\"https://bar.foo/atom\" >" +
		"<link href=\"https://bar.foo/atom\" type=\"application/atom+xml\" title=\"  Atom feed\" >" +
		"<link href=\"https://bar.foo/wp-content/atom\" >" +
		"<link href=\"https://bar.foo/rss\" type=\"application/rss+xml\" >" +
		"<link href=\"https://bar.foo/atom\" type=\"application/atom+xml\" >" +
		"</head>" +
		"<body></body>"
	var p = Parser{origURL: origURL, document: getDocument(html)}
	p.ShouldFindSyndicationSource()
	var d = p.Parse()

	if len(d.Feeds) != 1 {
		t.Errorf("Incorrect number of feeds: Wanted %d; got %d", 4, len(d.Feeds))
	}
	f := d.Feeds[0]
	u := origURL.String() + "/feed/"

	if f.URL.String() != u {
		t.Errorf("Incorrect feed URL: Wanted %s; got %s", u, f.URL.String())
	}
	if f.Type != syndication.RSS {
		t.Errorf("Incorrect document Lang: Wanted %s; got %s", syndication.RSS, f.Type)
	}
	if f.Title != defaultWPFeedTitle {
		t.Errorf("Incorrect document Lang: Wanted %s; got %s", defaultWPFeedTitle, f.Title)
	}
}
