package gql

import (
	"github/mickaelvieira/taipan/internal/gql/loaders"
	"github/mickaelvieira/taipan/internal/gql/resolvers"
	"github/mickaelvieira/taipan/internal/repository"
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

func mustParse(content string, r *repository.Repositories) *graphql.Schema {
	var dataloaders = &loaders.Loaders{Repositories: r}

	resolvers := resolvers.Resolvers{
		Dataloaders:  dataloaders,
		Repositories: r,
	}

	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	schema := graphql.MustParseSchema(content, &resolvers, opts...)

	return schema
}

// LoadAndParseSchema load and parse the graphQL schema
func LoadAndParseSchema(path string, r *repository.Repositories) *graphql.Schema {
	return mustParse(mustLoad(path), r)
}
