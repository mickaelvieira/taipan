package parser

import (
	"io"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// MakeURLAbs defines the type of the anonymous function that makes a URL absolute
type MakeURLAbs = func(*url.URL) *url.URL

// getAbsURLCreator returns a function that can append the base domain to a URL if it is not absolute
func getAbsURLCreator(b *url.URL) MakeURLAbs {
	return func(u *url.URL) *url.URL {
		if !u.IsAbs() {
			username := b.User.Username()
			password, _ := b.User.Password()
			if username != "" {
				if password != "" {
					u.User = url.UserPassword(username, password)
				} else {
					u.User = url.User(username)
				}
			}
			u.Scheme = b.Scheme
			u.Host = b.Hostname()
		}
		return u
	}
}

// Parse parses the html tree and creates our parsed document
func Parse(URL *url.URL, r io.Reader) (*Document, error) {
	document, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	makeAbs := getAbsURLCreator(URL)

	// we collect meta and link tags upfront
	var metaTags, linkTags []*goquery.Selection
	document.Find("meta").Each(func(i int, s *goquery.Selection) {
		metaTags = append(metaTags, s)
	})
	document.Find("link").Each(func(i int, s *goquery.Selection) {
		linkTags = append(linkTags, s)
	})

	var p = Parser{
		makeAbs:  makeAbs,
		metaTags: metaTags,
		linkTags: linkTags,
		document: document,
	}

	doc := &Document{
		origURL:   URL,
		title:     p.Title(),
		desc:      p.Description(),
		Lang:      p.Lang(),
		Charset:   p.Charset(),
		canonical: p.CanonicalURL(),
		twitter:   p.TwitterTags(),
		facebook:  p.FacebookTags(),
		Feeds:     p.Feeds(),
	}

	return doc, nil
}
