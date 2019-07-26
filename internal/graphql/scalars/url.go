package scalars

import (
	"encoding/json"
	"fmt"

	"github/mickaelvieira/taipan/internal/domain/url"
)

// URL is a custom GraphQL type to represent a URL.
type URL struct {
	url *url.URL
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (URL) ImplementsGraphQLType(name string) bool {
	return name == "URL"
}

// UnmarshalGraphQL is a custom unmarshaler for URL
func (s *URL) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		var err error
		s.url, err = url.FromRawURL(input)
		return err
	default:
		return fmt.Errorf("URL only accept string type as input")
	}
}

// MarshalJSON is a custom marshaler for URL
func (s URL) MarshalJSON() ([]byte, error) {
	if s.url != nil {
		return json.Marshal(s.url.String())
	}
	return json.Marshal(nil)
}

// ToDomain returns the domain URL
func (s *URL) ToDomain() *url.URL {
	return s.url
}

// NewURL wrapped domain URL into a scalar URL
func NewURL(u *url.URL) URL {
	return URL{url: u}
}
