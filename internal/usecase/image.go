package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/s3"
)

// HandleImage this usecase:
// - fetches the document image
// - uploads it to AWS S3
// - Updates the DB
func HandleImage(ctx context.Context, d *document.Document, repositories *repository.Repositories) error {
	if d.Image == nil {
		fmt.Println("Document does not have an image associated")
		return nil
	}

	if d.Image.Name != "" {
		fmt.Printf("Image has already been fetched with name %s\n", d.Image.Name)
		return nil
	}

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

	return nil
}
