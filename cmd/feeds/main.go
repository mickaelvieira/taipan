package main

import (
	"context"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/bookmark"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/usecase"
	"io"
	"log"
	"net/url"

	"github/mickaelvieira/taipan/internal/app"

	"github.com/mmcdole/gofeed"
)

func main() {
	app.LoadEnvironment()

	fp := gofeed.NewParser()
	ctx := context.Background()
	repositories := repository.GetRepositories()

	feeds, err := repositories.Feeds.GetNewFeeds(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, feed := range feeds {
		var URL *url.URL
		var prevResult *client.Result
		var reader io.Reader
		var result *client.Result

		prevResult, err = repositories.Botlogs.FindLatestByURI(ctx, feed.URL)
		URL, err = url.ParseRequestURI(feed.URL)

		cl := client.Client{}
		result, reader, err = cl.Fetch(URL)

		log.Println(feed)

		if err == nil {
			if !result.IsContentDifferent(prevResult) {
				var content *gofeed.Feed
				content, err = fp.Parse(reader)
				if err == nil {
					for _, item := range content.Items {
						log.Println(item.Link)
						var bookmark *bookmark.Bookmark
						bookmark, err = usecase.Bookmark(ctx, item.Link, repositories)

						if err != nil {
							log.Println(err)
						} else {
							log.Println(bookmark)
						}
					}
				}
			}
		}
	}
}
