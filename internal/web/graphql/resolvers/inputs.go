package resolvers

type CursorPaginationInput struct {
	From  *string
	To    *string
	Limit *int32
}

type OffsetPaginationInput struct {
	Offset *int32
	Limit  *int32
}

type SubscriptionSearchInput struct {
	Terms []string
	Tags  []string
}

type BookmarkSearchInput struct {
	Terms []string
}

type DocumentSearchInput struct {
	Terms []string
}

type SyndicationSearchInput struct {
	Terms       []string
	Tags        []string
	ShowDeleted bool
	PausedOnly  bool
}

type UserInput struct {
	Firstname string
	Lastname  string
	Image     string
}
