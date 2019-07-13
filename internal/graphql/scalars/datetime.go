package scalars

import (
	"encoding/json"
	"fmt"
	"time"
)

// Datetime is a custom GraphQL type to represent an instant in time.
type Datetime struct {
	w time.Time
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (Datetime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

// UnmarshalGraphQL is a custom unmarshaler for DateTime
func (t *Datetime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		var err error
		t.w, err = time.Parse(time.RFC3339, input)
		return err
	default:
		return fmt.Errorf("Datetime only accept string type as input")
	}
}

// MarshalJSON is a custom marshaler for Time
func (t Datetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.w.Format(time.RFC3339))
}

// NewDatetime wrapped domain Time into a scalar Datetime
func NewDatetime(t time.Time) Datetime {
	return Datetime{w: t}
}
