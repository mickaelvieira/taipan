package main

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/repository"
	"log"

	"github/mickaelvieira/taipan/internal/app"

	"github.com/mmcdole/gofeed"
)

func main() {
	app.LoadEnvironment()

	fp := gofeed.NewParser()

	repo := repository.GetRepositories().Feeds

	ctx := context.Background()
	feeds, err := repo.GetNewFeeds(ctx)

	if err != nil {
		log.Fatal(err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.URL)
		feed, err := fp.ParseURL(feed.URL)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(feed.Title)
		for _, item := range feed.Items {
			fmt.Println(item.Link)
		}
	}
}
