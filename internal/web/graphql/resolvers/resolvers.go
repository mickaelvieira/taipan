package resolvers

import (
	"github/mickaelvieira/taipan/internal/web/graphql/generated"
)

type RootResolver struct{}

func (r *RootResolver) Mutation() generated.MutationResolver         {}
func (r *RootResolver) Query() generated.QueryResolver               {}
func (r *RootResolver) Subscription() generated.SubscriptionResolver {}
