package repository

import (
	"github/mickaelvieira/taipan/internal/db"
	"strings"
)

// Scanable sql.Rows or sql.Row
type Scanable interface {
	Scan(...interface{}) error
}

// Repositories holds a reference to the repositories
type Repositories struct {
	Users       *UserRepository
	Syndication *SyndicationRepository
	Documents   *DocumentRepository
	Bookmarks   *BookmarkRepository
	Botlogs     *BotlogRepository
}

// GetRepositories builds the repository holder
func GetRepositories() *Repositories {
	var db = db.GetDB()

	return &Repositories{
		Users:       &UserRepository{db: db},
		Syndication: &SyndicationRepository{db: db},
		Documents:   &DocumentRepository{db: db},
		Bookmarks:   &BookmarkRepository{db: db},
		Botlogs:     &BotlogRepository{db: db},
	}
}

func formatQuery(query string) string {
	return strings.ReplaceAll(strings.ReplaceAll(query, "\t", " "), "\n", "")
}
