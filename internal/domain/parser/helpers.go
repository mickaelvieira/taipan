package parser

import (
	"github/mickaelvieira/taipan/internal/domain/feed"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func IsURLValid(URL string) bool {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}
	return true
}

func NormalizeAttr(val string) string {
	return strings.ToLower(strings.Trim(val, " "))
}

func NormalizeHTMLText(val string) string {
	return strings.Trim(val, " \n")
}

func ParseCharset(metaTags []*goquery.Selection) string {
	var charset string // @TODO I might give it a default value

	// try <meta charset="" >
	for _, s := range metaTags {
		var exist bool
		charset, exist = s.Attr("charset")
		if exist {
			charset = NormalizeAttr(charset)
			break
		}
	}

	if charset == "" {
		// otherwise <meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1">
		for _, s := range metaTags {
			he := NormalizeAttr(s.AttrOr("http-equiv", ""))
			ct := NormalizeAttr(s.AttrOr("content", ""))

			if he == "content-type" && strings.Contains(ct, "charset") {
				var c = "charset="
				var i = strings.LastIndex(ct, "charset")
				charset = ct[i+len(c) : len(ct)]
			}
		}
	}

	return charset
}

func ParseTitle(doc *goquery.Document, metaTags []*goquery.Selection) string {
	var title string

	doc.Find("title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title = s.Text()
		return false
	})

	if title == "" {
		for _, s := range metaTags {
			name, exist := s.Attr("name")
			name = NormalizeAttr(name)
			if exist && name == "title" {
				title = s.AttrOr("content", "")
				break
			}
		}
	}
	return NormalizeHTMLText(title)
}

func ParseLang(doc *goquery.Document, metaTags []*goquery.Selection) string {
	var lang string

	doc.Find("html").EachWithBreak(func(i int, s *goquery.Selection) bool {
		lang = NormalizeAttr(s.AttrOr("lang", ""))
		return false
	})

	if lang == "" {
		for _, s := range metaTags {
			he := NormalizeAttr(s.AttrOr("http-equiv", ""))
			if he == "content-language" {
				lang = NormalizeAttr(s.AttrOr("content", ""))
			}
		}
	}
	return lang
}

func ParseDescription(metaTags []*goquery.Selection) string {
	var desc string
	for _, s := range metaTags {
		name, exist := s.Attr("name")
		name = NormalizeAttr(name)
		if exist && name == "description" {
			desc = s.AttrOr("content", "")
			break
		}
	}
	return NormalizeHTMLText(desc)
}

func ParseCanonicalURL(linkTags []*goquery.Selection) string {
	var URL string
	for _, s := range linkTags {
		rel, exist := s.Attr("href")
		rel = NormalizeAttr(rel)
		if exist && rel == "canonical" {
			URL = NormalizeAttr(s.AttrOr("href", ""))
			break
		}
	}
	return URL
}

func ParseFeeds(linkTags []*goquery.Selection) []*feed.Feed {
	var feeds []*feed.Feed
	for _, s := range linkTags {
		h := NormalizeAttr(s.AttrOr("href", ""))
		t := NormalizeAttr(s.AttrOr("title", ""))
		ft, err := feed.GetFeedType(NormalizeAttr(s.AttrOr("type", "")))

		if IsURLValid(h) && err == nil {
			feed := feed.Feed{Type: ft, URL: h, Title: t}
			feeds = append(feeds, &feed)
		}
	}
	return feeds
}
