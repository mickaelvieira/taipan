package url

import (
	"database/sql/driver"
	"errors"
	"fmt"
	neturl "net/url"
	"strings"
)

// URL represents a URL within the application
type URL struct {
	*neturl.URL
}

// RemoveGAParams removes Google Analytics parameters
func (url *URL) RemoveGAParams() {
	params := strings.Split(url.RawQuery, "&")
	var p []string
	for _, param := range params {
		s := strings.Split(param, "=")
		if strings.Index(s[0], "utm_") == -1 {
			p = append(p, param)
		}
	}
	url.RawQuery = strings.Join(p, "&")
}

// Value converts the value going into the DB
func (url *URL) Value() (driver.Value, error) {
	value, err := neturl.QueryUnescape(url.String())
	return value, err
}

// Scan converts the value coming from the DB
func (url *URL) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	if v, ok := value.([]byte); ok {
		var u *neturl.URL
		u, err := neturl.ParseRequestURI(string(v))
		if err != nil {
			return errors.New("failed to parse URL during scanning")
		}
		*url = URL{u}
		return nil
	}
	return errors.New("failed to scan URL")
}

func removeFragment(rawURL string) string {
	var i = strings.LastIndex(rawURL, "#")
	if i < 0 {
		return rawURL
	}
	return rawURL[0:i]
}

// FromRawURL returns an URL struct only when the raw URL is absolute. It also removes the URL fragment
func FromRawURL(rawURL string) (*URL, error) {
	u, err := neturl.ParseRequestURI(removeFragment(rawURL))
	if err != nil || !u.IsAbs() {
		return nil, fmt.Errorf("Invalid URL '%s'", rawURL)
	}

	n := &URL{
		URL: u,
	}

	n.RemoveGAParams()

	return n, nil
}
