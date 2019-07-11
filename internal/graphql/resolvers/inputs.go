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

type searchSourcesInput struct {
	IsPaused bool
}

type userInput struct {
	Firstname string
	Lastname  string
	Image     string
}
