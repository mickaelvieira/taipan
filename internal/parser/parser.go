package parser

import (
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/feed"
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/uri"
	"html"
	"net/url"
	"strings"

	strip "github.com/grokify/html-strip-tags-go"

	"github.com/PuerkitoBio/goquery"
)

type social struct {
	Title       string
	Description string
	Image       *url.URL // @TODO we need to handle base64 encoded images
	URL         *url.URL
}

// Parser parses the document
type Parser struct {
	origURL   *url.URL
	metaTags  []*goquery.Selection
	linkTags  []*goquery.Selection
	document  *goquery.Document
	docTitle  string
	docDesc   string
	canonical *url.URL
	twitter   *social
	facebook  *social
}

// Parse parse the document
func (p *Parser) Parse() *document.Document {
	// cache relevant tags
	p.document.Find("meta").Each(func(i int, s *goquery.Selection) {
		p.metaTags = append(p.metaTags, s)
	})
	p.document.Find("link").Each(func(i int, s *goquery.Selection) {
		p.linkTags = append(p.linkTags, s)
	})

	// gather some data upfront in order to get relevant information
	isWP := p.isWordpress()
	p.docTitle = p.parseTitle()
	p.docDesc = p.parseDescription()
	p.canonical = p.parseCanonicalURL()
	p.twitter = p.parseTwitterTags()
	p.facebook = p.parseFacebookTags()

	// get the data we need
	var url = p.url()
	var lang = p.parseLang()
	var charset = p.parseCharset()
	var title = p.title()
	var description = p.description()
	var image = p.image()

	var feeds []*feed.Feed
	if isWP {
		// WP is a bit silly in the way it handles RSS feeds
		// Usually only the comments feed is available on the page
		// So we construct the default feed URL ourselves
		feeds = p.getWordpressFeed()
	} else {
		// @TODO I need to include a black list to avoid adding unwanted feeds such as github commits
		feeds = p.parseFeeds()
	}

	return document.New(
		url,
		lang,
		charset,
		title,
		description,
		image,
		feeds,
	)
}

func (p *Parser) isWordpress() bool {
	var isWP bool
	for _, s := range p.linkTags {
		rel, exist := s.Attr("rel")
		url := p.normalizeAttrValue(s.AttrOr("href", ""))
		rel = p.normalizeAttrValue(rel)
		if exist && rel == "stylesheet" && strings.Contains(url, "wp-content") {
			isWP = true
			break
		}
	}
	return isWP
}

func (p *Parser) getWordpressFeed() []*feed.Feed {
	var feeds []*feed.Feed
	url := &url.URL{Path: "/feed/"} // default WP feed
	url = p.makeURLAbs(url)
	feed := feed.New(
		&uri.URI{URL: url},
		"wordpress feed",
		feed.RSS,
	)
	feeds = append(feeds, feed)
	return feeds
}

// URL retrieves the URL of the document. It will first try to grab the canonical URL, if there isn't one
// it will try to get one from the social media tags. If it can find any it will return the URL provided by the user
func (p *Parser) url() *uri.URI {
	var du *url.URL
	if p.canonical != nil {
		du = p.canonical
	} else if p.facebook.URL != nil {
		du = p.facebook.URL
	} else if p.twitter.URL != nil {
		du = p.twitter.URL
	} else {
		du = p.removeFragment(p.origURL)
	}
	return &uri.URI{URL: du}
}

// Title retrieve the title of the document. If there isn't a title tag,
// it will try to get the title from the socual media tags
func (p *Parser) title() string {
	var t string
	if p.docTitle != "" {
		t = p.docTitle
	} else if p.facebook.Title != "" {
		t = p.facebook.Title
	} else if p.twitter.Title != "" {
		t = p.twitter.Title
	}
	return t
}

// Description retrieve the description of the document. If there isn't a description meta tag,
// it will try to get the description from the socual media tags
func (p *Parser) description() string {
	var de string
	if p.docDesc != "" {
		de = p.docDesc
	} else if p.facebook.Description != "" {
		de = p.facebook.Description
	} else if p.twitter.Description != "" {
		de = p.twitter.Description
	}
	return de
}

// Image retrieves the image URL from the social media tag. It will return an empty string if there isn't any
func (p *Parser) image() *image.Image {
	var iu *url.URL
	if p.facebook.Image != nil {
		iu = p.facebook.Image
	} else if p.twitter.Image != nil {
		iu = p.twitter.Image
	}

	if iu == nil {
		return nil
	}

	return &image.Image{
		URL: &uri.URI{URL: iu},
	}
}

func (p *Parser) parseCharset() string {
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

func (p *Parser) parseTitle() string {
	var title string
	p.document.Find("title").EachWithBreak(func(i int, s *goquery.Selection) bool {
		title = strip.StripTags(s.Text())
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

func (p *Parser) parseLang() string {
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

func (p *Parser) parseDescription() string {
	var desc string
	for _, s := range p.metaTags {
		name, exist := s.Attr("name")
		name = p.normalizeAttrValue(name)
		if exist && name == "description" {
			desc = strip.StripTags(s.AttrOr("content", ""))
			break
		}
	}
	return p.normalizeHTMLText(desc)
}

func (p *Parser) parseCanonicalURL() *url.URL {
	var rawURL string
	for _, s := range p.linkTags {
		rel, exist := s.Attr("rel")
		rel = p.normalizeAttrValue(rel)
		if exist && rel == "canonical" {
			rawURL = s.AttrOr("href", "")
			break
		}
	}
	return p.parseAndNormalizeRawURL(rawURL)
}

func (p *Parser) parseFeeds() []*feed.Feed {
	var feeds []*feed.Feed
	for _, s := range p.linkTags {
		url := p.normalizeAttrValue(s.AttrOr("href", ""))
		title := p.normalizeAttrValue(p.normalizeHTMLText(s.AttrOr("title", "")))
		feedType, err := feed.GetFeedType(p.normalizeAttrValue(s.AttrOr("type", "")))
		urlFeed := p.parseAndNormalizeRawURL(url)
		if err == nil && urlFeed != nil {
			feed := feed.New(
				&uri.URI{URL: urlFeed},
				title,
				feedType,
			)
			feeds = append(feeds, feed)
		}
	}
	return feeds
}

func (p *Parser) parseFacebookTags() *social {
	return p.parseSocialTags("og:", "property")
}

func (p *Parser) parseTwitterTags() *social {
	return p.parseSocialTags("twitter:", "name")
}

func (p *Parser) parseSocialTags(prefix string, property string) *social {
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
				title = p.normalizeHTMLText(val)
			case "description":
				desc = p.normalizeHTMLText(val)
			case "image":
				image = val
			case "image:src":
				image = val
			case "url":
				url = val
			}
		}
	}

	return &social{
		Title:       title,
		Description: desc,
		Image:       p.parseAndNormalizeRawURL(image),
		URL:         p.parseAndNormalizeRawURL(url),
	}
}

func (p *Parser) normalizeAttrValue(val string) string {
	return strings.ToLower(strings.Trim(val, " "))
}

func (p *Parser) normalizeHTMLText(val string) string {
	return html.UnescapeString(strings.Trim(val, " \n"))
}

func (p *Parser) makeURLAbs(u *url.URL) *url.URL {
	if !u.IsAbs() {
		username := p.origURL.User.Username()
		password, _ := p.origURL.User.Password()
		if username != "" {
			if password != "" {
				u.User = url.UserPassword(username, password)
			} else {
				u.User = url.User(username)
			}
		}
		u.Scheme = p.origURL.Scheme
		u.Host = p.origURL.Hostname()
	}
	return u
}

func (p *Parser) parseAndNormalizeRawURL(rawURL string) *url.URL {
	URL, _ := url.ParseRequestURI(rawURL)
	if URL != nil {
		p.makeURLAbs(URL)
		p.removeFragment(URL)
	}
	return URL
}

func (p *Parser) removeFragment(u *url.URL) *url.URL {
	u.Fragment = ""
	return u
}
