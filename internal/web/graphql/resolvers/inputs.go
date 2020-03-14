package resolvers

// CursorPaginationInput cursor based pagination
type CursorPaginationInput struct {
	From  *string
	To    *string
	Limit *int32
}

// OffsetPaginationInput offset based pagination
type OffsetPaginationInput struct {
	Offset *int32
	Limit  *int32
}

// SortingInput sorting parameters
type SortingInput struct {
	By  *string
	Dir *string
}

// SubscriptionSearchInput user's subscriptions search input
type SubscriptionSearchInput struct {
	Terms []string
	Tags  []string
}

// BookmarkSearchInput user's bookmarks search input
type BookmarkSearchInput struct {
	Terms []string
}

// DocumentSearchInput document search input
type DocumentSearchInput struct {
	Terms []string
}

// SyndicationSearchInput syndication sources search input
type SyndicationSearchInput struct {
	Terms  []string
	Tags   []string
	Hidden bool
	Paused bool
	Sort   *SortingInput
}

// UserInput user profile input values
type UserInput struct {
	Firstname string
	Lastname  string
	Image     string
}
