package resolvers

import (
	"context"
	"github/mickaelvieira/taipan/internal/publisher"
	"github/mickaelvieira/taipan/internal/usecase"
	"github/mickaelvieira/taipan/internal/web/auth"
	"github/mickaelvieira/taipan/internal/web/clientid"
	"github/mickaelvieira/taipan/internal/web/graphql/generated"
)

// UsersQuery bookmarks' root resolver
type UsersQuery struct {
	repositories *repository.Repositories
	publisher    *publisher.Subscription
}

type UsersMutation struct {
	// repositories *repository.Repositories
	// publisher    *publisher.Subscription
}

// Update resolves the mutation
func (r *UsersMutation) Update(ctx context.Context, args struct {
	User generated.UserInput
}) (*generated.User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

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

	return resolve(r.repositories).user(user), nil
}

// Password resolves the mutation
func (r *UsersMutation) Password(ctx context.Context, args struct {
	Old string
	New string
}) (bool, error) {
	user := auth.FromContext(ctx)
	// clientID := clientid.FromContext(ctx)

	err := usecase.ChangePassword(
		ctx,
		r.repositories,
		user,
		args.Old,
		args.New,
	)
	if err != nil {
		return false, err
	}

	// r.publisher.Publish(
	// 	publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	// )

	return true, nil
}

// Theme resolves the mutation
func (r *UsersMutation) Theme(ctx context.Context, args struct {
	Theme string
}) (*generated.User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

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

	return resolve(r.repositories).user(user), nil
}

// CreateEmail --
func (r *UsersMutation) CreateEmail(ctx context.Context, args struct {
	Email string
}) (*generated.User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	err := usecase.CreateUserEmail(ctx, r.repositories, user, args.Email)
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}

// DeleteEmail --
func (r *UsersMutation) DeleteEmail(ctx context.Context, args struct {
	Email string
}) (*generated.User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	err := usecase.DeleteUserEmail(ctx, r.repositories, user, args.Email)
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}

// PrimaryEmail --
func (r *UsersMutation) PrimaryEmail(ctx context.Context, args struct {
	Email string
}) (*generated.User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	err := usecase.PrimaryUserEmail(ctx, r.repositories, user, args.Email)
	if err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}
