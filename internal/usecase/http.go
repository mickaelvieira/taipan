package usecase

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
	"log"
)

// FetchResource fetches the related resource
func FetchResource(ctx context.Context, repos *repository.Repositories, u *url.URL) (r *http.Result, err error) {
	c := http.Client{}

	if r := c.Get(u); r != nil {
		log.Printf("%v", r.RespHeaders)
		err = repos.Botlogs.Insert(ctx, r)
		if err != nil {
			return nil, err
		}
	}

	return r, c.Err()
}
