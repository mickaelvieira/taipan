package resolvers

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/auth"
	"github/mickaelvieira/taipan/internal/clientid"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/publisher"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"log"

	gql "github.com/graph-gophers/graphql-go"
)

// UsersResolver bookmarks' root resolver
type UsersResolver struct {
	repositories *repository.Repositories
	publisher    *publisher.Subscription
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

// Theme resolves the Theme field
func (r *UserResolver) Theme() string {
	return r.User.Theme
}

// Image resolves the Image field
func (r *UserResolver) Image() *UserImageResolver {
	if !r.User.HasImage() {
		return nil
	}

	return &UserImageResolver{
		Image: r.User.Image,
	}
}

// UserEventResolver resolves an bookmark event
type UserEventResolver struct {
	event *publisher.Event
}

// Item returns the event's message
func (r *UserEventResolver) Item() *UserResolver {
	u, ok := r.event.Payload.(*user.User)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a user")
	}
	return &UserResolver{User: u}
}

// Emitter returns the event's emitter ID
func (r *UserEventResolver) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *UserEventResolver) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *UserEventResolver) Action() string {
	return string(r.event.Action)
}

// LoggedIn resolves the query
func (r *UsersResolver) LoggedIn(ctx context.Context) (*UserResolver, error) {
	user := auth.FromContext(ctx)

	res := UserResolver{User: user}

	return &res, nil
}

// UserChanged subscribes to user event
func (r *RootResolver) UserChanged(ctx context.Context) <-chan *UserEventResolver {
	c := make(chan *UserEventResolver)
	s := &userSubscriber{events: c}
	r.publisher.Subscribe(publisher.TopicUser, s, ctx.Done())
	return c
}

// Update resolves the mutation
func (r *UsersResolver) Update(ctx context.Context, args struct {
	ID   string
	User userInput
}) (*UserResolver, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)
	if args.ID != user.ID {
		return nil, fmt.Errorf("You are not allowed to modify this user")
	}

	err := usecase.UpdateUser(
		ctx,
		r.repositories,
		user,
		args.User.Firstname,
		args.User.Lastname,
		args.User.Image,
	)
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	res := UserResolver{User: user}

	return &res, nil
}

// Theme resolves the mutation
func (r *UsersResolver) Theme(ctx context.Context, args struct {
	ID    string
	Theme string
}) (*UserResolver, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)
	if args.ID != user.ID {
		return nil, fmt.Errorf("You are not allowed to modify this user")
	}

	err := usecase.UpdateTheme(
		ctx,
		r.repositories,
		user,
		args.Theme,
	)
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	res := UserResolver{User: user}

	return &res, nil
}
