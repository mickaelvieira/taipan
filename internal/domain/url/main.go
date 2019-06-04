package url

import (
	"database/sql/driver"
	"errors"
	neturl "net/url"
	"strings"
)

// URL represents a URL within the application
type URL struct {
	*neturl.URL
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
		return nil, errors.New("Invalid URL")
	}
	return &URL{URL: u}, nil
}
