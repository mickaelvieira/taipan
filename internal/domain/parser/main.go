package parser

import (
	"errors"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func FetchAndParse(URL string) (*bookmark.Bookmark, error) {
	if !IsURLValid(URL) {
		return nil, errors.New("Invalid URL")
	}

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var metaTags, linkTags []*goquery.Selection

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		metaTags = append(metaTags, s)
	})

	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		linkTags = append(linkTags, s)
	})

	var image, url string
	var title = ParseTitle(doc, metaTags)
	var lang = ParseLang(doc, metaTags)
	var desc = ParseDescription(metaTags)
	var charset = ParseCharset(metaTags)
	var canonical = ParseCanonicalURL(linkTags)
	var tw = ParseTwitterTags(metaTags)
	var fb = ParseFacebookTags(metaTags)
	// var feeds = ParseFeeds(linkTags)

	// Try to get a good URL
	if IsURLValid(canonical) {
		url = canonical
	} else if IsURLValid(fb.URL) {
		url = fb.URL
	} else if IsURLValid(tw.URL) {
		url = tw.URL
	} else {
		url = URL
	}

	// Try to get a good title
	if title == "" && fb.Title != "" {
		title = fb.Title
	}

	if title == "" && tw.Title != "" {
		title = tw.Title
	}

	// Try to get a good description
	if desc == "" && fb.Description != "" {
		desc = fb.Description
	}

	if desc == "" && tw.Description != "" {
		desc = tw.Description
	}

	// Try to get a good image
	if image == "" && fb.Image != "" {
		image = fb.Image
	}

	if image == "" && tw.Image != "" {
		image = tw.Image
	}

	bookmark := bookmark.New(
		url,
		lang,
		charset,
		title,
		desc,
		image,
	)

	return bookmark, nil
}
