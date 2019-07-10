package resolvers

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"

	gql "github.com/graph-gophers/graphql-go"
)

// UsersResolver bookmarks' root resolver
type UsersResolver struct {
	repositories *repository.Repositories
}

// UserResolver resolves the user entity
type UserResolver struct {
	*user.User
}

// ID resolves the ID field
func (r *UserResolver) ID() gql.ID {
	return gql.ID(r.User.ID)
}

// Username resolves the Username field
func (r *UserResolver) Username() string {
	return r.User.Username
}

// Firstname resolves the Firstname field
func (r *UserResolver) Firstname() string {
	return r.User.Firstname
}

// Lastname resolves the Lastname field
func (r *UserResolver) Lastname() string {
	return r.User.Lastname
}

// Image resolves the Image field
func (r *UserResolver) Image() *UserImageResolver {
	if r.User.Image == nil || r.User.Image.Name == "" {
		return nil
	}

	return &UserImageResolver{
		Image: r.User.Image,
	}
}

// LoggedIn resolves the query
func (r *UsersResolver) LoggedIn(ctx context.Context) (*UserResolver, error) {
	user := auth.FromContext(ctx)

	res := UserResolver{User: user}

	return &res, nil
}

// UserInput data received from graphQL
type userInput struct {
	Firstname string
	Lastname  string
	Image     string
}

// Update resolves the mutation
func (r *UsersResolver) Update(ctx context.Context, args struct {
	ID   string
	User userInput
}) (*UserResolver, error) {
	u := auth.FromContext(ctx)
	if args.ID != u.ID {
		return nil, fmt.Errorf("You are not allowed to modify this user")
	}

	user, err := usecase.UpdateUser(
		ctx,
		u,
		args.User.Firstname,
		args.User.Lastname,
		args.User.Image,
		r.repositories,
	)
	if err != nil {
		return nil, err
	}

	res := UserResolver{User: user}

	return &res, nil
}
