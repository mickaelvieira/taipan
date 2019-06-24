package scalars

import (
	"encoding/json"
	"fmt"
	"time"
)

// DateTime is a custom GraphQL type to represent an instant in time.
type DateTime struct {
	time.Time
}

// ImplementsGraphQLType maps this custom Go type
// to the graphql scalar type in the schema.
func (DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

// UnmarshalGraphQL is a custom unmarshaler for DateTime
func (t *DateTime) UnmarshalGraphQL(input interface{}) error {
	switch input := input.(type) {
	case string:
		var err error
		t.Time, err = time.Parse(time.RFC3339, input)
		return err
	default:
		return fmt.Errorf("Datetime only accept string type as input")
	}
}

// MarshalJSON is a custom marshaler for Time
func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format(time.RFC3339))
}
