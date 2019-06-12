package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/repository"
	"github/mickaelvieira/taipan/internal/s3"
	"time"
)

// HandleImage this usecase:
// - fetches the document image
// - uploads it to AWS S3
// - Updates the DB
func HandleImage(ctx context.Context, d *document.Document, repositories *repository.Repositories) (err error) {
	if d.Image == nil {
		fmt.Println("Document does not have an image associated")
		return
	}

	if d.Image.Name != "" {
		fmt.Printf("Image has already been fetched with name %s\n", d.Image.Name)
		return
	}

	result, err := FetchResource(ctx, d.Image.URL, repositories)
	if err != nil {
		return
	}

	err = s3.Upload(d.Image, result, result.Content)
	if err != nil {
		return
	}

	// Image was uploaded at the point so we can update the document
	d.UpdatedAt = time.Now()

	err = repositories.Documents.UpdateImage(ctx, d)
	return
}
