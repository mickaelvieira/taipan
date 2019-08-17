package resolvers

type cursorPaginationInput struct {
	From  *string
	To    *string
	Limit *int32
}

type offsetPaginationInput struct {
	Offset *int32
	Limit  *int32
}

type subscriptionSearchInput struct {
	Terms       []string
	ShowDeleted bool
	PausedOnly  bool
}

type bookmarkSearchInput struct {
	Terms []string
}

type documentSearchInput struct {
	Terms []string
}

type searchSourcesInput struct {
	IsPaused bool
}

type userInput struct {
	Firstname string
	Lastname  string
	Image     string
}
