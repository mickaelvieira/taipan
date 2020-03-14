package graphql

import (
	"github.com/mickaelvieira/taipan/internal/repository"
	"github.com/mickaelvieira/taipan/internal/web/graphql/resolvers"
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

func mustParse(content string, repositories *repository.Repositories) *graphql.Schema {
	resolvers := resolvers.GetRootResolver(repositories)
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(content, resolvers, opts...)

	return schema
}

// LoadAndParseSchema load and parse the graphQL schema
func LoadAndParseSchema(path string, r *repository.Repositories) *graphql.Schema {
	return mustParse(mustLoad(path), r)
}
