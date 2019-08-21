package resolvers

type cursorPaginationInput struct {
	From  *string
	To    *string
	Limit *int32
}

type OffsetPaginationInput struct {
	Offset *int32
	Limit  *int32
}

type SubscriptionSearchInput struct {
	Terms       []string
	ShowDeleted bool
	PausedOnly  bool
}

type BookmarkSearchInput struct {
	Terms []string
}

type DocumentSearchInput struct {
	Terms []string
}

type SearchSourcesInput struct {
	IsPaused bool
}

type UserInput struct {
	Firstname string
	Lastname  string
	Image     string
}
