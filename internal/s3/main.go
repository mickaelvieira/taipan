package s3

import (
	"bytes"
	"fmt"
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/image"
	img "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	_ "golang.org/x/image/webp"
)

// Format image formats available
type Format struct {
	GIF  string
	JPEG string
	PNG  string
	WEBP string
}

// ContentType content types available
type ContentType struct {
	GIF  string
	JPG  string
	JPEG string
	PNG  string
	WEBP string
}

var contentType = &ContentType{
	GIF:  "image/gif",
	JPG:  "image/jpg",
	JPEG: "image/jpeg",
	PNG:  "image/png",
	WEBP: "image/webp",
}

var format = &Format{
	GIF:  "gif",
	JPEG: "jpeg",
	PNG:  "png",
	WEBP: "webp",
}

func getFormatFromContentType(i string) string {
	var o string
	switch i {
	case contentType.JPG:
	case contentType.JPEG:
		o = format.JPEG
	case contentType.GIF:
		o = format.GIF
	case contentType.PNG:
		o = format.PNG
	case contentType.WEBP:
		o = format.WEBP
	}
	return o
}

func getContentTypeFromFormat(i string) string {
	var o string
	switch i {
	case format.JPEG:
		o = contentType.JPEG
	case format.GIF:
		o = contentType.GIF
	case format.PNG:
		o = contentType.PNG
	case format.WEBP:
		o = contentType.WEBP
	}

	return o
}

// Upload uploads a file to the S3 bucket
func Upload(i *image.Image) error {
	bucket := os.Getenv("AWS_BUCKET")

	resp, err := http.Get(i.URL.String())
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")

	// Get file configuration
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	filename := checksum.FromBytes(b).String()
	reader := bytes.NewReader(b)
	config, format, err := img.DecodeConfig(reader)

	if err != nil {
		if err != img.ErrFormat {
			return err
		}
		format = ""
	}

	if format == "" && contentType != "" {
		format = getFormatFromContentType(contentType)
	}

	if contentType == "" && format != "" {
		contentType = getContentTypeFromFormat(format)
	}

	// image's filename
	if format != "" {
		filename = filename + "." + format
	}

	i.Name = filename
	i.Width = int32(config.Width)
	i.Height = int32(config.Height)
	i.Format = format

	// Recreate a IO reader from buffered bytes
	reader = bytes.NewReader(b)

	// Upload file to AWS S3
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	cacheControl := "public, max-age=2592000"

	output, err := uploader.Upload(&s3manager.UploadInput{
		ACL:          aws.String("public-read"),
		Bucket:       aws.String(bucket),
		Key:          aws.String(i.Name),
		Body:         reader,
		ContentType:  &contentType,
		CacheControl: &cacheControl,
	})

	if err != nil {
		return err
	}

	fmt.Printf("Uploaded %s \n", output.Location)

	return nil
}
