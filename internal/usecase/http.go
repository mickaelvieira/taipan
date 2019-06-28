package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
)

// FetchResource fetches the related resource
func FetchResource(ctx context.Context, u *url.URL, repositories *repository.Repositories) (r *http.Result, err error) {
	http := client.Client{}
	r, err = http.Get(u)
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
