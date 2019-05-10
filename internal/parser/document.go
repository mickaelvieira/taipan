package parser

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/domain/types"
	"net/url"
)

// Document is the data structure representing a parsed document
type Document struct {
	origURL   *url.URL
	title     string
	desc      string
	canonical *url.URL
	twitter   *Social
	facebook  *Social
	Lang      string
	Charset   string
	Canonical string
	Feeds     []*feed.Feed
}

// URL retrieves the URL of the document. It will first try to grab the canonical URL, if there isn't one
// it will try to get one from the social media tags. If it can find any it will return the URL provided by the user
func (d *Document) URL() *types.URI {
	var du *url.URL
	if d.canonical != nil {
		du = d.canonical
	} else if d.facebook.URL != nil {
		du = d.facebook.URL
	} else if d.twitter.URL != nil {
		du = d.twitter.URL
	} else {
		du = d.origURL
	}
	return &types.URI{URL: du}
}

// Title retrieve the title of the document. If there isn't a title tag,
// it will try to get the title from the socual media tags
func (d *Document) Title() string {
	var t string
	if d.title != "" {
		t = d.title
	} else if d.facebook.Title != "" {
		t = d.facebook.Title
	} else if d.twitter.Title != "" {
		t = d.twitter.Title
	}
	return t
}

// Description retrieve the description of the document. If there isn't a description meta tag,
// it will try to get the description from the socual media tags
func (d *Document) Description() string {
	var de string
	if d.desc != "" {
		de = d.desc
	} else if d.facebook.Description != "" {
		de = d.facebook.Description
	} else if d.twitter.Description != "" {
		de = d.twitter.Description
	}
	return de
}

// Image retrieves the image URL from the social media tag. It will return an empty string if there isn't any
func (d *Document) Image() *bookmark.Image {
	var iu *url.URL
	if d.facebook.Image != nil {
		iu = d.facebook.Image
	} else if d.twitter.Image != nil {
		iu = d.twitter.Image
	}

	if iu == nil {
		return nil
	}

	return &bookmark.Image{
		URL: &types.URI{URL: iu},
	}
}

// ToBookmark is a factory to create a bookmark entity from the document
func (d *Document) ToBookmark() *bookmark.Bookmark {
	return bookmark.New(
		d.URL(),
		d.Lang,
		d.Charset,
		d.Title(),
		d.Description(),
		d.Image(),
	)
}
