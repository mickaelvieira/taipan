package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
)

// FetchResource fetches the related resource
func FetchResource(ctx context.Context, u *url.URL, repositories *repository.Repositories) (r *client.Result, err error) {
	// @TODO that might be nice to do a HEAD request
	// to get the last modified date before fetching the entire document
	http := client.Client{}
	r, err = http.Fetch(u)
	if err != nil {
		return
	}

	// Store the result of HTTP request
	err = repositories.Botlogs.Insert(ctx, r)
	if err != nil {
		return
	}

	// We only want successful requests
	if r.RespStatusCode != 200 {
		err = fmt.Errorf("Unable to fetch the document: %s", r.RespReasonPhrase)
	}

	return
}
