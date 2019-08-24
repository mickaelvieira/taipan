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

type SortingInput struct {
	By  *string
	Dir *string
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
	Terms  []string
	Tags   []string
	Hidden bool
	Paused bool
	Sort   *SortingInput
}

type UserInput struct {
	Firstname string
	Lastname  string
	Image     string
}
