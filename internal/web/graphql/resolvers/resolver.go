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
	r  *repository.Repositories
	sl *dataloader.Loader
	ll *dataloader.Loader
}

func (f *resolver) bookmarks(results []*bookmark.Bookmark) []*BookmarkResolver {
	bookmarks := make([]*BookmarkResolver, len(results))
	for i, d := range results {
		bookmarks[i] = f.bookmark(d)
	}
	return bookmarks
}

func (f *resolver) bookmark(b *bookmark.Bookmark) *BookmarkResolver {
	return &BookmarkResolver{
		b:  b,
		r:  f.r,
		sl: f.sl,
		ll: f.ll,
	}
}

func (f *resolver) documents(results []*document.Document) []*DocumentResolver {
	documents := make([]*DocumentResolver, len(results))
	for i, d := range results {
		documents[i] = f.document(d)
	}
	return documents
}

func (f *resolver) document(d *document.Document) *DocumentResolver {
	return &DocumentResolver{
		d:  d,
		r:  f.r,
		sl: f.sl,
		ll: f.ll,
	}
}

func (f *resolver) sources(results []*syndication.Source) []*SourceResolver {
	sources := make([]*SourceResolver, len(results))
	for i, d := range results {
		sources[i] = f.source(d)
	}
	return sources
}

func (f *resolver) source(s *syndication.Source) *SourceResolver {
	return &SourceResolver{
		s:  s,
		r:  f.r,
		ll: f.ll,
	}
}

func (f *resolver) subscriptions(results []*subscription.Subscription) []*SubscriptionResolver {
	subscription := make([]*SubscriptionResolver, len(results))
	for i, d := range results {
		subscription[i] = f.subscription(d)
	}
	return subscription
}

func (f *resolver) subscription(s *subscription.Subscription) *SubscriptionResolver {
	return &SubscriptionResolver{
		s:  s,
		r:  f.r,
		ll: f.ll,
	}
}

func (f *resolver) logs(results []*http.Result) []*LogResolver {
	logs := make([]*LogResolver, len(results))
	for i, d := range results {
		logs[i] = f.log(d)
	}
	return logs
}

func (f *resolver) log(l *http.Result) *LogResolver {
	return &LogResolver{
		l: l,
		r: f.r,
	}
}

func resolve(r *repository.Repositories) *resolver {
	return &resolver{
		r:  r,
		sl: loaders.GetSource(r.Syndication),
		ll: loaders.GetLogs(r.Botlogs),
	}
}
