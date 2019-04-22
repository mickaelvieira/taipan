package parser

import (
	"errors"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Social struct {
	Title       string
	Description string
	Image       string
	URL         string
}

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

	var title = ParseTitle(doc, metaTags)
	var lang = ParseLang(doc, metaTags)
	var desc = ParseDescription(metaTags)
	var charset = ParseCharset(metaTags)
	var canonical = ParseCanonicalURL(linkTags)
	// var feeds = ParseFeeds(linkTags)

	bookmark := &bookmark.Bookmark{
		URL:         URL,
		Lang:        lang,
		Title:       title,
		Description: desc,
		Charset:     charset,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Status:      bookmark.FETCHED,
	}

	log.Println(bookmark)
	log.Println(canonical)
	// log.Println(feeds)

	return bookmark, nil
}
