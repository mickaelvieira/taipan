package scalars

import (
	"encoding/json"
	"fmt"

	"github/mickaelvieira/taipan/internal/domain/url"
)

// URL is a custom GraphQL type to represent a URL.
type URL struct {
	*url.URL
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (URL) ImplementsGraphQLType(name string) bool {
	return name == "URL"
}

// UnmarshalGraphQL is a custom unmarshaler for URL
func (u *URL) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		var err error
		u.URL, err = url.FromRawURL(input)
		return err
	default:
		return fmt.Errorf("URL only accept string type as input")
	}
}

// MarshalJSON is a custom marshaler for URL
func (u URL) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.URL.String())
}
