package usecase

import (
	"context"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/document"
	"github/mickaelvieira/taipan/internal/domain/image"
	"github/mickaelvieira/taipan/internal/domain/user"
	"github/mickaelvieira/taipan/internal/logger"
	"github/mickaelvieira/taipan/internal/repository"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UpdateToS3 uploads the image to S3 bucket
func UpdateToS3(name string, contentType string, r io.Reader) error {
	logger.Info("Upload to S3")
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

	logger.Info(fmt.Sprintf("Uploaded %s ", output.Location))

	return nil
}

// HandleImage this usecase:
// - fetches the document image
// - uploads it to AWS S3
// - Updates the DB
func HandleImage(ctx context.Context, repos *repository.Repositories, d *document.Document) error {
	if d.Image == nil {
		logger.Info("Document does not have an image associated")
		return nil
	}

	if d.Image.Name != "" {
		logger.Info(fmt.Sprintf("Image has already been fetched with name %s", d.Image.Name))
		return nil
	}

	result, err := FetchResource(ctx, repos, d.Image.URL)
	if err != nil {
		return err
	}

	if !result.RequestWasSuccessful() {
		return fmt.Errorf(result.GetFailureReason())
	}

	d.Image.Name = image.GetName(result.Checksum, result.ContentType)
	d.Image.Format = image.GetExtension(result.ContentType)

	dm, r := image.GetDimensions(result.Content)
	d.Image.SetDimensions(dm.Width, dm.Height)

	if err := UpdateToS3(d.Image.Name, result.ContentType, r); err != nil {
		return err
	}

	// Image was uploaded at the point so we can update the document
	d.UpdatedAt = time.Now()

	if err := repos.Documents.UpdateImage(ctx, d); err != nil {
		return err
	}

	return nil
}

// HandleAvatar this usecase:
// - retrieves image's information from base 64 data
// - uploads it to AWS S3
// - Updates the DB
func HandleAvatar(ctx context.Context, repos *repository.Repositories, usr *user.User, s string) error {
	c := image.GetContentType(s)
	d := image.GetBase64Data(s)
	r := image.GetBase64Reader(d)
	// l := image.GetBase64DataLen(d)

	// b := bytes.NewBuffer(make([]byte, 0, l))
	// b.ReadFrom(r)

	var cs checksum.Checksum
	cs, r = checksum.FromReader(r)

	i := &user.Image{
		Name:   image.GetName(cs, c),
		Format: image.GetExtension(c),
	}

	if usr.Image != nil && usr.Image.Name == i.Name {
		return nil
	}

	dm, r := image.GetDimensions(r)
	i.SetDimensions(dm.Width, dm.Height)

	if err := UpdateToS3(i.Name, c, r); err != nil {
		return err
	}

	usr.Image = i
	usr.UpdatedAt = time.Now()

	if err := repos.Users.UpdateImage(ctx, usr); err != nil {
		return err
	}

	return nil
}
