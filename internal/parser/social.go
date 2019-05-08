package parser

import "net/url"

// Social represents the data from social media tags
type Social struct {
	Title       string
	Description string
	Image       *url.URL
	URL         *url.URL
}
