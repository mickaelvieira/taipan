package parser

import (
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/domain/types"
	"html"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Parser parses the document
type Parser struct {
	makeAbs  MakeURLAbs
	metaTags []*goquery.Selection
	linkTags []*goquery.Selection
	document *goquery.Document
}

// Charset parses page's charset
func (p *Parser) Charset() string {
	var charset string
	for _, s := range p.metaTags {
		var exist bool
		charset, exist = s.Attr("charset")
		if exist {
			charset = p.normalizeAttrValue(charset)
			break
		}
	}
	if charset == "" {
		for _, s := range p.metaTags {
			he := p.normalizeAttrValue(s.AttrOr("http-equiv", ""))
			ct := p.normalizeAttrValue(s.AttrOr("content", ""))
			if he == "content-type" && strings.Contains(ct, "charset") {
				var c = "charset="
				var i = strings.LastIndex(ct, "charset")
				charset = ct[i+len(c) : len(ct)]
			}
		}
	}
	return charset
}

// Title parses document's title
func (p *Parser) Title() string {
	var title string
	p.document.Find("title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title = s.Text()
		return false
	})
	if title == "" {
		for _, s := range p.metaTags {
			name, exist := s.Attr("name")
			name = p.normalizeAttrValue(name)
			if exist && name == "title" {
				title = s.AttrOr("content", "")
				break
			}
		}
	}
	return p.normalizeHTMLText(title)
}

// Lang parses document's language
func (p *Parser) Lang() string {
	var lang string

	p.document.Find("html").EachWithBreak(func(i int, s *goquery.Selection) bool {
		lang = p.normalizeAttrValue(s.AttrOr("lang", ""))
		return false
	})

	if lang == "" {
		for _, s := range p.metaTags {
			he := p.normalizeAttrValue(s.AttrOr("http-equiv", ""))
			if he == "content-language" {
				lang = p.normalizeAttrValue(s.AttrOr("content", ""))
			}
		}
	}
	return lang
}

// Description parses document's description
func (p *Parser) Description() string {
	var desc string
	for _, s := range p.metaTags {
		name, exist := s.Attr("name")
		name = p.normalizeAttrValue(name)
		if exist && name == "description" {
			desc = s.AttrOr("content", "")
			break
		}
	}
	return p.normalizeHTMLText(desc)
}

// CanonicalURL parses document canonical URL
func (p *Parser) CanonicalURL() *url.URL {
	var rawURL string
	for _, s := range p.linkTags {
		rel, exist := s.Attr("href")
		rel = p.normalizeAttrValue(rel)
		if exist && rel == "canonical" {
			rawURL = s.AttrOr("href", "")
			break
		}
	}
	return p.parseAndNormalizeRawURL(rawURL)
}

// Feeds parses feeds links
func (p *Parser) Feeds() []*feed.Feed {
	var feeds []*feed.Feed
	for _, s := range p.linkTags {
		url := p.normalizeAttrValue(s.AttrOr("href", ""))
		title := p.normalizeAttrValue(s.AttrOr("title", ""))
		feedType, err := feed.GetFeedType(p.normalizeAttrValue(s.AttrOr("type", "")))
		urlFeed := p.parseAndNormalizeRawURL(url)
		if err == nil && urlFeed != nil {
			feed := feed.New(
				&types.URI{URL: urlFeed},
				title,
				feedType,
			)
			feeds = append(feeds, &feed)
		}
	}
	return feeds
}

// FacebookTags parses Facebook meta tags
func (p *Parser) FacebookTags() *Social {
	return p.parseSocialTags("og:", "property")
}

// TwitterTags parses Twitter meta tags
func (p *Parser) TwitterTags() *Social {
	return p.parseSocialTags("twitter:", "name")
}

func (p *Parser) parseSocialTags(prefix string, property string) *Social {
	var title, desc, image, url string
	for _, s := range p.metaTags {
		var exist bool
		var prop string
		prop, exist = s.Attr(property)
		val := s.AttrOr("content", "")
		if exist && strings.HasPrefix(prop, prefix) && val != "" {
			prop = strings.ToLower(prop[len(prefix):len(prop)])
			switch prop {
			case "title":
				title = val
			case "description":
				desc = val
			case "image":
				image = val
			case "url":
				url = val
			}
		}
	}

	return &Social{
		Title:       title,
		Description: desc,
		Image:       p.parseAndNormalizeRawURL(image),
		URL:         p.parseAndNormalizeRawURL(url),
	}
}

// NormalizeAttrValue trims whitespaces and transfor the string to lower case
func (p *Parser) normalizeAttrValue(val string) string {
	return strings.ToLower(strings.Trim(val, " "))
}

// NormalizeHTMLText trims break line and whitespaces
func (p *Parser) normalizeHTMLText(val string) string {
	return html.UnescapeString(strings.Trim(val, " \n"))
}

func (p *Parser) parseAndNormalizeRawURL(rawURL string) *url.URL {
	URL, _ := url.ParseRequestURI(rawURL)
	if URL != nil {
		p.makeAbs(URL)
		removeFragment(URL)
	}
	return URL
}
