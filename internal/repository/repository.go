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
	Users           *UserRepository
	Emails          *UserEmailRepository
	EmailsConfirm   *UserEmailConfirmRepository
	NewsFeed        *NewsFeedRepository
	Syndication     *SyndicationRepository
	SyndicationTags *SyndicationTagsRepository
	Subscriptions   *SubscriptionRepository
	Documents       *DocumentRepository
	Bookmarks       *BookmarkRepository
	Botlogs         *BotlogRepository
	PasswordReset   *PasswordResetRepository
}

// GetRepositories builds the repository holder
func GetRepositories() *Repositories {
	var db = db.GetDB()

	return &Repositories{
		Users:           &UserRepository{db: db},
		Emails:          &UserEmailRepository{db: db},
		EmailsConfirm:   &UserEmailConfirmRepository{db: db},
		NewsFeed:        &NewsFeedRepository{db: db},
		Syndication:     &SyndicationRepository{db: db},
		SyndicationTags: &SyndicationTagsRepository{db: db},
		Subscriptions:   &SubscriptionRepository{db: db},
		Documents:       &DocumentRepository{db: db},
		Bookmarks:       &BookmarkRepository{db: db},
		Botlogs:         &BotlogRepository{db: db},
		PasswordReset:   &PasswordResetRepository{db: db},
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

func getDocumentSearch(terms []string) (string, []interface{}) {
	var s string
	var a []interface{}
	if len(terms) > 0 {
		// @TODO the WITH QUERY EXPANSION mode is awesoe but is quite slow,
		// we need to see find out how we can improve that
		s = "AND MATCH(d.title, d.description) AGAINST(? IN NATURAL LANGUAGE MODE)"
		a = append(a, strings.Join(terms, " "))
	}
	return s, a
}

func getSyndicationSearch(terms []string) (string, []interface{}) {
	var s string
	var a []interface{}

	if len(terms) > 0 {
		var or []string
		for _, t := range terms {
			or = append(or, "s.url LIKE ?")
			a = append(a, "%"+t+"%")
		}

		or = append(or, "MATCH(s.title) AGAINST(? IN NATURAL LANGUAGE MODE)")

		a = append(a, strings.Join(terms, " "))
		s = fmt.Sprintf("(%s)", strings.Join(or, " OR "))
	}

	return s, a
}
