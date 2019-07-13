package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/repository"
	"io"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UpdateToS3 uploads the image to S3 bucket
func UpdateToS3(name string, contentType string, r io.Reader) error {
	log.Println("Upload to S3")
	bucket := os.Getenv("AWS_BUCKET")

	// Upload file to AWS S3
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	cacheControl := "public, max-age=" + os.Getenv("AWS_MAX_AGE")

	output, err := uploader.Upload(&s3manager.UploadInput{
		ACL:          aws.String("public-read"),
		Bucket:       aws.String(bucket),
		Key:          aws.String(name),
		Body:         r,
		ContentType:  &contentType,
		CacheControl: &cacheControl,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Uploaded %s \n", output.Location)

	return nil
}

// HandleImage this usecase:
// - fetches the document image
// - uploads it to AWS S3
// - Updates the DB
func HandleImage(ctx context.Context, repos *repository.Repositories, d *document.Document) (err error) {
	if d.Image == nil {
		fmt.Println("Document does not have an image associated")
		return
	}

	if d.Image.Name != "" {
		fmt.Printf("Image has already been fetched with name %s\n", d.Image.Name)
		return
	}

	result, err := FetchResource(ctx, repos, d.Image.URL)
	if err != nil {
		return
	}

	d.Image.Name = image.GetName(result.Checksum, result.ContentType)
	d.Image.Format = image.GetExtension(result.ContentType)

	dm, r := image.GetDimensions(result.Content)
	if err != nil {
		return
	}

	d.Image.SetDimensions(dm.Width, dm.Height)

	err = UpdateToS3(d.Image.Name, result.ContentType, r)
	if err != nil {
		return
	}

	// Image was uploaded at the point so we can update the document
	d.UpdatedAt = time.Now()

	err = repos.Documents.UpdateImage(ctx, d)
	return
}

// HandleAvatar this usecase:
// - retrieves image's information from base 64 data
// - uploads it to AWS S3
// - Updates the DB
func HandleAvatar(ctx context.Context, repos *repository.Repositories, usr *user.User, s string) (err error) {
	c := image.GetContentType(s)
	d := image.GetBase64Data(s)
	r := image.GetBase64Reader(d)

	var cs checksum.Checksum
	cs, r = checksum.FromReader(r)

	i := &user.Image{
		Name:   image.GetName(cs, c),
		Format: image.GetExtension(c),
	}

	if usr.Image != nil && usr.Image.Name == i.Name {
		return
	}

	var dm *image.Dimensions
	dm, r = image.GetDimensions(r)
	if err != nil {
		return
	}

	i.SetDimensions(dm.Width, dm.Height)

	err = UpdateToS3(i.Name, c, r)
	if err != nil {
		return
	}

	usr.Image = i
	usr.UpdatedAt = time.Now()

	err = repos.Users.UpdateImage(ctx, usr)

	return err
}
