package helpers

import "net/url"

// RemoveFragment removes the fragment from the URL
func RemoveFragment(u *url.URL) *url.URL {
	u.Fragment = ""

	return u
}
