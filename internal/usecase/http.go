package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/url"
	"github/mickaelvieira/taipan/internal/repository"
	"io"
)

// FetchResource fetches the related resource
func FetchResource(ctx context.Context, u *url.URL, repositories *repository.Repositories) (*client.Result, io.Reader, error) {
	// @TODO that might be nice to do a HEAD request
	// to get the last modified date before fetching the entire document
	http := client.Client{}
	result, reader, err := http.Fetch(u)
	if err != nil {
		return nil, nil, err
	}

	// Store the result of HTTP request
	err = repositories.Botlogs.Insert(ctx, result)
	if err != nil {
		return nil, nil, err
	}

	// We only want successful requests
	if result.RespStatusCode != 200 {
		return result, nil, fmt.Errorf("Unable to fetch the document: %s", result.RespReasonPhrase)
	}

	return result, reader, err
}
