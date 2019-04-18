package gql

import (
	"github/mickaelvieira/taipan/internal/gql/resolvers"
	"io/ioutil"
	"log"

	graphql "github.com/graph-gophers/graphql-go"
)

func mustLoad(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func mustParse(content string) *graphql.Schema {
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(content, &resolvers.Resolvers{}, opts...)
	return schema
}

// LoadAndParseSchema load and parse the graphQL schema
func LoadAndParseSchema(path string) *graphql.Schema {
	return mustParse(mustLoad(path))
}
