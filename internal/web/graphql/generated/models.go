// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

type Node interface {
	IsNode()
}

type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AppQuery struct {
	Info *AppInfo `json:"info"`
}

type BookmarkSearchInput struct {
	Terms []string `json:"terms"`
}

type CursorPaginationInput struct {
	From  *string `json:"from"`
	To    *string `json:"to"`
	Limit *int    `json:"limit"`
}

type DocumentSearchInput struct {
	Terms []string `json:"terms"`
}

type Email struct {
	ID          string  `json:"id"`
	Value       string  `json:"value"`
	IsPrimary   bool    `json:"isPrimary"`
	IsConfirmed bool    `json:"isConfirmed"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	ConfirmedAt *string `json:"confirmedAt"`
}

func (Email) IsNode() {}

type Image struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Format string `json:"format"`
}

type OffsetPaginationInput struct {
	Offset *int `json:"offset"`
	Limit  *int `json:"limit"`
}

type SearchSourcesInput struct {
	IsPaused bool `json:"isPaused"`
}

type SubscriptionSearchInput struct {
	Terms       []string `json:"terms"`
	ShowDeleted bool     `json:"showDeleted"`
	PausedOnly  bool     `json:"pausedOnly"`
}

type User struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Emails    []*Email   `json:"emails"`
	Theme     string     `json:"theme"`
	Image     *Image     `json:"image"`
	Stats     *UserStats `json:"stats"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
}

func (User) IsNode() {}

type UserInput struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Image     string `json:"image"`
}

type UserStats struct {
	ID            string `json:"id"`
	Bookmarks     int    `json:"bookmarks"`
	Favorites     int    `json:"favorites"`
	ReadingList   int    `json:"readingList"`
	Subscriptions int    `json:"subscriptions"`
}

func (UserStats) IsNode() {}
