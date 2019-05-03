package url

import (
	"log"
	"net/url"
)

// GetAbsURLCreator creates a function to make URL absolute
func GetAbsURLCreator(b *url.URL) func(*url.URL) *url.URL {
	return func(u *url.URL) *url.URL {
		if !u.IsAbs() {
			username := b.User.Username()
			password, _ := b.User.Password()
			if username != "" {
				if password != "" {
					u.User = url.UserPassword(username, password)
				} else {
					u.User = url.User(username)
				}
			}

			u.Scheme = b.Scheme
			u.Host = b.Hostname()
		}
		return u
	}
}

//MakeAbs makes the URL absolute if it is not
func MakeAbs(docURL *url.URL, rawURL string) *url.URL {
	URL, err := url.ParseRequestURI(rawURL)

	if err != nil {
		return &url.URL{}
	}

	if !URL.IsAbs() {
		username := docURL.User.Username()
		password, _ := docURL.User.Password()
		if username != "" {
			if password != "" {
				URL.User = url.UserPassword(username, password)
			} else {
				URL.User = url.User(username)
			}
		}

		URL.Scheme = docURL.Scheme
		URL.Host = docURL.Hostname()
	}

	return URL
}

// Unescape unespace query string
func Unescape(e string) string {
	u, _ := url.QueryUnescape(e)
	return u
}
