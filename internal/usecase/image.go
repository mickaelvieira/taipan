package usecase

import (
	"context"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/s3"
)

// HandleImage this usecase:
// - fetches the document image
// - uploads it to AWS S3
// - Updates the DB
func HandleImage(ctx context.Context, d *document.Document, repositories *repository.Repositories) error {
	// An image was found in the document
	if d.Image != nil {
		result, reader, err := FetchResource(ctx, d.Image.URL, repositories)
		if err != nil {
			return err
		}

		err = s3.Upload(d.Image, result, reader)
		if err != nil {
			return err
		}

		err = repositories.Documents.UpdateImage(ctx, d)
		if err != nil {
			return err
		}
	}

	return nil
}
