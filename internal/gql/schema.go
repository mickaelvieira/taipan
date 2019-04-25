package gql

import (
	"github/mickaelvieira/taipan/internal/gql/loaders"
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
	resolvers := resolvers.Resolvers{Dataloaders: &loaders.Loaders{}}
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(content, &resolvers, opts...)

	return schema
}

// LoadAndParseSchema load and parse the graphQL schema
func LoadAndParseSchema(path string) *graphql.Schema {
	return mustParse(mustLoad(path))
}
