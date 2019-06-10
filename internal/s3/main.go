package s3

import (
	"fmt"
	"github/mickaelvieira/taipan/internal/client"
	"github/mickaelvieira/taipan/internal/domain/image"
	img "image"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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

func getFormatFromContentType(i string) (o string) {
	switch i {
	case contentType.JPG:
		o = format.JPEG
	case contentType.JPEG:
		o = format.JPEG
	case contentType.GIF:
		o = format.GIF
	case contentType.PNG:
		o = format.PNG
	case contentType.WEBP:
		o = format.WEBP
	}
	return
}

func getContentTypeFromFormat(i string) (o string) {
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
	return
}

// func resizeImage(ct string, in io.Reader) (out io.Reader, err error) {
// 	var i img.Image
// 	switch ct {
// 	case contentType.JPG:
// 		i, err = jpeg.Decode(in)
// 	case contentType.JPEG:
// 		i, err = jpeg.Decode(in)
// 	case contentType.GIF:
// 		i, err = gif.Decode(in)
// 	case contentType.PNG:
// 		i, err = png.Decode(in)
// 	default:
// 		err = fmt.Errorf("Cannot decode type %s", ct)
// 	}

// 	if err != nil {
// 		log.Printf("Decode %s", err)
// 		return
// 	}

// 	i = imaging.Resize(i, 800, 0, imaging.Lanczos)

// 	var b bytes.Buffer
// 	w := bufio.NewWriter(&b)

// 	switch ct {
// 	case contentType.JPG:
// 		err = jpeg.Encode(w, i, &jpeg.Options{Quality: 100})
// 	case contentType.JPEG:
// 		err = jpeg.Encode(w, i, &jpeg.Options{Quality: 100})
// 	case contentType.GIF:
// 		err = gif.Encode(w, i, &gif.Options{})
// 	case contentType.PNG:
// 		err = png.Encode(w, i)
// 	default:
// 		err = fmt.Errorf("Cannot encode type %s", ct)
// 	}

// 	if err != nil {
// 		log.Printf("Encode %s", err)
// 		return
// 	}

// 	out = bytes.NewReader(b.Bytes())

// 	log.Println("Image was resized")

// 	return
// }

// Upload uploads a file to the S3 bucket
func Upload(i *image.Image, result *client.Result, reader io.Reader) error {
	bucket := os.Getenv("AWS_BUCKET")

	contentType := result.ContentType
	filename := result.Checksum.String()

	// @TODO DO I need to update the checksum?
	// var err error
	// reader, err = resizeImage(contentType, reader)
	// if err != nil {
	// 	return err
	// }

	// Read image configuration from IO
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

	// if contentType == "" && format != "" {
	// 	contentType = getContentTypeFromFormat(format)
	// }

	// log.Printf("Content type %s", contentType)
	// log.Printf("Format %s", format)
	// log.Printf("Width %d", config.Width)
	// log.Printf("Height %d", config.Height)

	// append image's extension
	if format != "" {
		filename = filename + "." + format
	}

	i.Name = filename
	i.Width = int32(config.Width)
	i.Height = int32(config.Height)
	i.Format = format

	// Upload file to AWS S3
	sess := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(sess)
	cacheControl := "public, max-age=" + os.Getenv("AWS_MAX_AGE")

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
