package graphql

import (
	"io/ioutil"
	"log"

	graphql "github.com/graph-gophers/graphql-go"
)

// LoadAndParseSchema load and parse the graphQL schema
func LoadAndParseSchema() *graphql.Schema {
	content, err := ioutil.ReadFile("./schema.graphql")

	if err != nil {
		log.Fatal(err)
	}

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(string(content), &Resolvers{}, opts...)

	return schema
}
