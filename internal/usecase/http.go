package usecase

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/http"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
)

// FetchResource fetches the related resource
func FetchResource(ctx context.Context, repos *repository.Repositories, u *url.URL) (*http.Result, error) {
	c := http.Client{}
	r := c.Get(u)
	if err := repos.Botlogs.Insert(ctx, r); err != nil {
		return r, err
	}
	return r, nil
}
