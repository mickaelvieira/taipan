package resolvers

// BookmarkCollectionResolver resolver
type BookmarkCollectionResolver struct {
	Results *[]*BookmarkResolver
	Total   int32
	Offset  int32
	Limit   int32
}

// UserBookmarkCollectionResolver resolver
type UserBookmarkCollectionResolver struct {
	Results *[]*UserBookmarkResolver
	Total   int32
	Offset  int32
	Limit   int32
}
