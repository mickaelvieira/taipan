package repository

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/db"
	"strings"
)

// Scanable sql.Rows or sql.Row
type Scanable interface {
	Scan(...interface{}) error
}

// Repositories holds a reference to the repositories
type Repositories struct {
	Users         *UserRepository
	Emails        *UserEmailRepository
	NewsFeed      *NewsFeedRepository
	Syndication   *SyndicationRepository
	Subscriptions *SubscriptionRepository
	Documents     *DocumentRepository
	Bookmarks     *BookmarkRepository
	Botlogs       *BotlogRepository
}

// GetRepositories builds the repository holder
func GetRepositories() *Repositories {
	var db = db.GetDB()

	return &Repositories{
		Users:         &UserRepository{db: db},
		Emails:        &UserEmailRepository{db: db},
		NewsFeed:      &NewsFeedRepository{db: db},
		Syndication:   &SyndicationRepository{db: db},
		Subscriptions: &SubscriptionRepository{db: db},
		Documents:     &DocumentRepository{db: db},
		Bookmarks:     &BookmarkRepository{db: db},
		Botlogs:       &BotlogRepository{db: db},
	}
}

func formatQuery(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(query, "'", "`"), "\t", " "), "\n", "")
}

func getMultiInsertPlacements(t int, n int) string {
	a := make([]string, t)
	for i := range a {
		p := make([]string, n)
		for j := range p {
			p[j] = "?"
		}
		if len(p) > 0 {
			a[i] = fmt.Sprintf("(%s)", strings.Join(p, ", "))
		}
	}
	return strings.Join(a, ", ")
}
