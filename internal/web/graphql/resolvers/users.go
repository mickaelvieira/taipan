package resolvers

import (
	"context"
	"github.com/mickaelvieira/taipan/internal/aggregator"
	"github.com/mickaelvieira/taipan/internal/domain/user"
	"github.com/mickaelvieira/taipan/internal/publisher"
	"github.com/mickaelvieira/taipan/internal/repository"
	"github.com/mickaelvieira/taipan/internal/usecase"
	"github.com/mickaelvieira/taipan/internal/web/auth"
	"github.com/mickaelvieira/taipan/internal/web/clientid"
	"github.com/mickaelvieira/taipan/internal/web/graphql/loaders"
	"github.com/mickaelvieira/taipan/internal/web/graphql/scalars"
	"log"

	gql "github.com/graph-gophers/graphql-go"
)

// UserRootResolver bookmarks' root resolver
type UserRootResolver struct {
	repositories *repository.Repositories
	publisher    *publisher.Subscription
}

// User resolves the user entity
type User struct {
	user         *user.User
	repositories *repository.Repositories
}

// ID resolves the ID field
func (r *User) ID() gql.ID {
	return gql.ID(r.user.ID)
}

// Firstname resolves the Firstname field
func (r *User) Firstname() string {
	return r.user.Firstname
}

// Lastname resolves the Lastname field
func (r *User) Lastname() string {
	return r.user.Lastname
}

// UserStats resolves the user entity
type UserStats struct {
	user *user.User
}

func (r *UserStats) getTotal(ctx context.Context, ty aggregator.AggType) (int32, error) {
	var t int32

	l := loaders.FromContext(ctx)
	if l == nil {
		return t, ErrLoadersNotFound
	}

	d, err := l.UsersStats.Load(ctx, aggregator.NewLoaderKey(r.user, ty))()
	if err != nil {
		return t, err
	}

	var ok bool
	t, ok = d.(int32)
	if !ok {
		return t, ErrDataTypeIsNotValid
	}

	return t, nil
}

// Bookmarks count the number of bookmarks
func (r *UserStats) Bookmarks(ctx context.Context) (int32, error) {
	return r.getTotal(ctx, aggregator.Bookmarks)
}

// Favorites count the number of favorites
func (r *UserStats) Favorites(ctx context.Context) (int32, error) {
	return r.getTotal(ctx, aggregator.Favorites)
}

// ReadingList count the number of bookmarks in the reading list
func (r *UserStats) ReadingList(ctx context.Context) (int32, error) {
	return r.getTotal(ctx, aggregator.ReadingList)
}

// Subscriptions count the number of subscriptions
func (r *UserStats) Subscriptions(ctx context.Context) (int32, error) {
	return r.getTotal(ctx, aggregator.Subscriptions)
}

// ID resolves the ID field
func (r *UserStats) ID() gql.ID {
	return gql.ID(r.user.ID)
}

// Emails resolves the Emails field
func (r *User) Emails(ctx context.Context) ([]*Email, error) {
	l := loaders.FromContext(ctx)
	if l == nil {
		return nil, ErrLoadersNotFound
	}

	d, err := l.Emails.Load(ctx, r.user)()
	if err != nil {
		return nil, err
	}

	emails, ok := d.([]*user.Email)
	if !ok {
		return nil, ErrDataTypeIsNotValid
	}

	return resolve(r.repositories).emails(emails), nil
}

// Stats resolves the Stats field
func (r *User) Stats(ctx context.Context) *UserStats {
	return &UserStats{user: r.user}
}

// Theme resolves the Theme field
func (r *User) Theme() string {
	return r.user.Theme
}

// Image resolves the Image field
func (r *User) Image() *UserImage {
	if !r.user.HasImage() {
		return nil
	}

	return &UserImage{
		Image: r.user.Image,
	}
}

// CreatedAt resolves the CreatedAt field
func (r *User) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.user.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *User) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.user.UpdatedAt)
}

// Email --
type Email struct {
	email        *user.Email
	repositories *repository.Repositories
}

// ID resolves the ID field
func (r *Email) ID() gql.ID {
	return gql.ID(r.email.ID)
}

// Value resolves the Value field
func (r *Email) Value() string {
	return r.email.Value
}

// IsPrimary resolves the IsPrimary field
func (r *Email) IsPrimary() bool {
	return r.email.IsPrimary
}

// IsConfirmed resolves the IsConfirmed field
func (r *Email) IsConfirmed() bool {
	return r.email.IsConfirmed
}

// CreatedAt resolves the CreatedAt field
func (r *Email) CreatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.email.CreatedAt)
}

// UpdatedAt resolves the UpdatedAt field
func (r *Email) UpdatedAt() scalars.Datetime {
	return scalars.NewDatetime(r.email.UpdatedAt)
}

// ConfirmedAt resolves the UpdatedAt field
func (r *Email) ConfirmedAt() *scalars.Datetime {
	t := scalars.NewDatetime(r.email.ConfirmedAt)
	return &t
}

// UserEvent resolves an bookmark event
type UserEvent struct {
	event        *publisher.Event
	repositories *repository.Repositories
}

// Item returns the event's message
func (r *UserEvent) Item() *User {
	u, ok := r.event.Payload.(*user.User)
	if !ok {
		log.Fatal("Cannot resolve item, payload is not a user")
	}
	return resolve(r.repositories).user(u)
}

// Emitter returns the event's emitter ID
func (r *UserEvent) Emitter() string {
	return r.event.Emitter
}

// Topic returns the event's topic
func (r *UserEvent) Topic() string {
	return string(r.event.Topic)
}

// Action returns the event's action
func (r *UserEvent) Action() string {
	return string(r.event.Action)
}

type userSubscriber struct {
	repositories *repository.Repositories
	events       chan<- *UserEvent
}

func (s *userSubscriber) Publish(e *publisher.Event) {
	s.events <- &UserEvent{
		event:        e,
		repositories: s.repositories,
	}
}

// LoggedIn resolves the query
func (r *UserRootResolver) LoggedIn(ctx context.Context) (*User, error) {
	user := auth.FromContext(ctx)

	return resolve(r.repositories).user(user), nil
}

// UserChanged subscribes to user event
func (r *RootResolver) UserChanged(ctx context.Context) <-chan *UserEvent {
	// @TODO better handle authentication
	c := make(chan *UserEvent)
	s := &userSubscriber{
		events:       c,
		repositories: r.repositories,
	}
	r.publisher.Subscribe(publisher.TopicUser, s, ctx.Done())
	return c
}

// Update resolves the mutation
func (r *UserRootResolver) Update(ctx context.Context, a struct {
	User UserInput
}) (*User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	if err := usecase.UpdateUser(ctx, r.repositories, user, a.User.Firstname, a.User.Lastname, a.User.Image); err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}

// Password resolves the mutation
func (r *UserRootResolver) Password(ctx context.Context, a struct {
	Old string
	New string
}) (bool, error) {
	user := auth.FromContext(ctx)
	// clientID := clientid.FromContext(ctx)

	if err := usecase.ChangePassword(ctx, r.repositories, user, a.Old, a.New); err != nil {
		return false, err
	}

	// r.publisher.Publish(
	// 	publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	// )

	return true, nil
}

// Theme resolves the mutation
func (r *UserRootResolver) Theme(ctx context.Context, a struct {
	Theme string
}) (*User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	if err := usecase.UpdateTheme(ctx, r.repositories, user, a.Theme); err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}

// CreateEmail --
func (r *UserRootResolver) CreateEmail(ctx context.Context, a struct {
	Email string
}) (*User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	if err := usecase.CreateUserEmail(ctx, r.repositories, user, a.Email); err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}

// DeleteEmail --
func (r *UserRootResolver) DeleteEmail(ctx context.Context, a struct {
	Email string
}) (*User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	if err := usecase.DeleteUserEmail(ctx, r.repositories, user, a.Email); err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}

// PrimaryEmail --
func (r *UserRootResolver) PrimaryEmail(ctx context.Context, a struct {
	Email string
}) (*User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	if err := usecase.PrimaryUserEmail(ctx, r.repositories, user, a.Email); err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}

// SendConfirmationEmail --
func (r *UserRootResolver) SendConfirmationEmail(ctx context.Context, a struct {
	Email string
}) (*User, error) {
	user := auth.FromContext(ctx)
	clientID := clientid.FromContext(ctx)

	e, err := r.repositories.Emails.GetUserEmailByValue(ctx, user, a.Email)
	if err != nil {
		return nil, err
	}

	if err := usecase.CreateTokenAndSendConfirmationEmail(ctx, r.repositories, user, e); err != nil {
		return nil, err
	}

	r.publisher.Publish(
		publisher.NewEvent(clientID, publisher.TopicUser, publisher.Update, user),
	)

	return resolve(r.repositories).user(user), nil
}
