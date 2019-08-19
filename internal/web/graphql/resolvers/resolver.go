package resolvers

import (
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/subscription"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/web/graphql/loaders"

	"github.com/graph-gophers/dataloader"
)

type resolver struct {
	repositories *repository.Repositories
	sourceLoader *dataloader.Loader
	logLoader    *dataloader.Loader
}

func (r *resolver) bookmarks(results []*bookmark.Bookmark) []*Bookmark {
	bookmarks := make([]*Bookmark, len(results))
	for i, d := range results {
		bookmarks[i] = r.bookmark(d)
	}
	return bookmarks
}

func (r *resolver) bookmark(b *bookmark.Bookmark) *Bookmark {
	return &Bookmark{
		bookmark:     b,
		repositories: r.repositories,
		sourceLoader: r.sourceLoader,
		logLoader:    r.logLoader,
	}
}

func (r *resolver) documents(results []*document.Document) []*Document {
	documents := make([]*Document, len(results))
	for i, d := range results {
		documents[i] = r.document(d)
	}
	return documents
}

func (r *resolver) document(d *document.Document) *Document {
	return &Document{
		document:     d,
		repositories: r.repositories,
		sourceLoader: r.sourceLoader,
		logLoader:    r.logLoader,
	}
}

func (r *resolver) sources(results []*syndication.Source) []*Source {
	sources := make([]*Source, len(results))
	for i, d := range results {
		sources[i] = r.source(d)
	}
	return sources
}

func (r *resolver) source(s *syndication.Source) *Source {
	return &Source{
		source:     s,
		repository: r.repositories,
		logLoader:  r.logLoader,
	}
}

func (r *resolver) subscriptions(results []*subscription.Subscription) []*Subscription {
	subscription := make([]*Subscription, len(results))
	for i, d := range results {
		subscription[i] = r.subscription(d)
	}
	return subscription
}

func (r *resolver) subscription(s *subscription.Subscription) *Subscription {
	return &Subscription{
		subscription: s,
		repositories: r.repositories,
		logLoader:    r.logLoader,
	}
}

func (r *resolver) logs(results []*http.Result) []*Log {
	logs := make([]*Log, len(results))
	for i, d := range results {
		logs[i] = r.log(d)
	}
	return logs
}

func (r *resolver) log(l *http.Result) *Log {
	return &Log{
		l: l,
		r: r.repositories,
	}
}

func resolve(r *repository.Repositories) *resolver {
	return &resolver{
		repositories: r,
		sourceLoader: loaders.GetSource(r.Syndication),
		logLoader:    loaders.GetLogs(r.Botlogs),
	}
}
