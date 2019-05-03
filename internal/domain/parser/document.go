package parser

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"
)

// Social represents the data from social media tags
type Social struct {
	Title       string
	Description string
	Image       string
	URL         string
}

// Document is the data structure representing a parsed document
type Document struct {
	url         string
	title       string
	description string
	canonical   string
	Lang        string
	Charset     string
	Canonical   string
	twitter     *Social
	facebook    *Social
	Feeds       []*feed.Feed
}

// Title retrieve the title of the document. If there isn't a title tag,
// it will try to get the title from the socual media tags
func (d *Document) Title() string {
	if d.title != "" {
		return d.title
	}
	if d.facebook.Title != "" {
		return d.facebook.Title
	}
	if d.twitter.Title != "" {
		return d.twitter.Title
	}
	return ""
}

// Desc retrieve the description of the document. If there isn't a description meta tag,
// it will try to get the description from the socual media tags
func (d *Document) Desc() string {
	if d.description != "" {
		return d.description
	}
	if d.facebook.Description != "" {
		return d.facebook.Description
	}
	if d.twitter.Description != "" {
		return d.twitter.Description
	}
	return ""
}

// URL retrieves the URL of the document. It will first try to grab the canonical URL, if there isn't one
// it will try to get one from the social media tags. If it can find any it will return the URL provided by the user
func (d *Document) URL() string {
	if IsURLValid(d.canonical) {
		return d.canonical
	}
	if IsURLValid(d.facebook.URL) {
		return d.facebook.URL
	}
	if IsURLValid(d.twitter.URL) {
		return d.twitter.URL
	}
	return d.url
}

// Image retrieves the image URL from the social media tag. It will return an empty string if there isn't any
func (d *Document) Image() string {
	if d.facebook.Image != "" {
		return d.facebook.Image
	}
	if d.twitter.Image != "" {
		return d.twitter.Image
	}
	return ""
}

// ToBookmark is a factory to create a bookmark entity from the document
func (d *Document) ToBookmark() *bookmark.Bookmark {
	return bookmark.New(
		d.URL(),
		d.Lang,
		d.Charset,
		d.Title(),
		d.Desc(),
		d.Image(),
	)
}

// MustCreateDocument parses the html tree and creates our parsed document
func MustParse(r io.ReadCloser) *Document {
	document, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal(err)
	}

	var metaTags, linkTags []*goquery.Selection

	document.Find("meta").Each(func(i int, s *goquery.Selection) {
		metaTags = append(metaTags, s)
	})

	document.Find("link").Each(func(i int, s *goquery.Selection) {
		linkTags = append(linkTags, s)
	})

	return &Document{
		title:       ParseTitle(document, metaTags),
		description: ParseDescription(metaTags),
		Lang:        ParseLang(document, metaTags),
		Charset:     ParseCharset(metaTags),
		canonical:   ParseCanonicalURL(linkTags),
		twitter:     ParseTwitterTags(metaTags),
		facebook:    ParseFacebookTags(metaTags),
		Feeds:       ParseFeeds(linkTags),
	}
}
