package url

import (
	"database/sql/driver"
	"errors"
	"log"
	neturl "net/url"
	"strings"
)

// URL domain errors
var (
	ErrURLNotValid    = errors.New("URL is invalid")
	ErrURLNotAbsolute = errors.New("URL is not absolute")
)

// URL represents a URL within the application
type URL struct {
	*neturl.URL
}

// UnescapeString returns the URL as string but without being escaped
func (u *URL) UnescapeString() string {
	value, err := neturl.QueryUnescape(u.String())
	if err != nil {
		log.Fatal(err)
	}
	return value
}

// Value converts the value going into the DB
func (u *URL) Value() (driver.Value, error) {
	if u == nil {
		return nil, nil
	}
	value, err := neturl.QueryUnescape(u.String())
	return value, err
}

// Scan converts the value coming from the DB
func (u *URL) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if v, ok := value.([]byte); ok {
		p, err := FromRawURL(string(v))
		if err != nil {
			return errors.New("failed to parse URL during scanning")
		}
		*u = *p
		return nil
	}
	return errors.New("failed to scan URL")
}

func removeGAParams(u *neturl.URL) *neturl.URL {
	params := strings.Split(u.RawQuery, "&")
	var p []string
	for _, param := range params {
		s := strings.Split(param, "=")
		if strings.Index(s[0], "utm_") == -1 {
			p = append(p, param)
		}
	}
	u.RawQuery = strings.Join(p, "&")
	return u
}

func removeFragment(r string) string {
	var i = strings.LastIndex(r, "#")
	if i < 0 {
		return r
	}
	return r[0:i]
}

// FromRawURL returns an URL struct only when the raw URL is absolute. It also removes the URL fragment
func FromRawURL(r string) (*URL, error) {
	p, err := neturl.ParseRequestURI(removeFragment(r))
	if err != nil {
		return nil, ErrURLNotValid
	}

	p = removeGAParams(p)
	u := &URL{URL: p}

	if !p.IsAbs() {
		return u, ErrURLNotAbsolute
	}

	return u, nil
}
