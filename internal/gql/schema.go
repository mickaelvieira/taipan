package gql

import (
	"io/ioutil"
	"log"

	graphql "github.com/graph-gophers/graphql-go"
)

// LoadAndParseSchema load and parse the graphQL schema
func LoadAndParseSchema(path string) *graphql.Schema {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(string(content), &Resolvers{}, opts...)

	return schema
}
