package scalars

import (
	"encoding/json"
	"fmt"
	"time"
)

// Datetime is a custom GraphQL type to represent an instant in time.
type Datetime struct {
	datetime time.Time
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Datetime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

// UnmarshalGraphQL is a custom unmarshaler for DateTime
func (s *Datetime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		var err error
		s.datetime, err = time.Parse(time.RFC3339, input)
		return err
	default:
		return fmt.Errorf("Datetime only accept string type as input")
	}
}

// MarshalJSON is a custom marshaler for Time
func (s Datetime) MarshalJSON() ([]byte, error) {
	if !s.datetime.IsZero() {
		return json.Marshal(s.datetime.Format(time.RFC3339))
	}
	return json.Marshal(nil)
}

// NewDatetime wrapped domain Time into a scalar Datetime
func NewDatetime(t time.Time) Datetime {
	return Datetime{datetime: t}
}
