package document

import (
	"github/mickaelvieira/taipan/internal/domain/checksum"
	"github/mickaelvieira/taipan/internal/domain/syndication"
	"github/mickaelvieira/taipan/internal/domain/url"
	"time"
)

// Status defines the status of a bookmark
type Status string

// Status values
const (
	FETCHED  Status = "fetched"
	PENDING  Status = "pending"
	FETCHING Status = "fetching"
	FAILED   Status = "failed"
)

// Document struct represents a web document
type Document struct {
	ID          string
	Checksum    checksum.Checksum
	URL         *url.URL
	Lang        string
	Charset     string
	Title       string
	Description string
	Image       *Image
	Feeds       []*syndication.Source
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Deleted     bool
}

// HasImage determine whether the document has an image associated to it
func (d *Document) HasImage() bool {
	return d.Image != nil && d.Image.URL != nil && d.Image.Name != ""
}

func (d *Document) String() string {
	return d.ID
}

// Raw returns the key raws data
func (d *Document) Raw() interface{} {
	return d
}

// New creates a new document
func New(url *url.URL, lang string, charset string, title string, desc string, image *Image, feeds []*syndication.Source) *Document {
	return &Document{
		URL:         url,
		Lang:        lang,
		Charset:     charset,
		Title:       title,
		Description: desc,
		Image:       image,
		Feeds:       feeds,
		Status:      FETCHED,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// Image represents a bookmark's image
type Image struct {
	Name   string
	URL    *url.URL
	Width  int32
	Height int32
	Format string
}

// SetDimensions image's information
func (i *Image) SetDimensions(w int, h int) {
	i.Width = int32(w)
	i.Height = int32(h)
}

func (i *Image) String() string {
	if i.URL == nil {
		return ""
	}
	return i.URL.String()
}

// NewImage returns a document's image
func NewImage(rawURL string, name string, width int32, height int32, format string) (*Image, error) {
	URL, err := url.FromRawURL(rawURL)
	if err != nil {
		return nil, err
	}

	var i = Image{
		URL:    URL,
		Name:   name,
		Width:  width,
		Height: height,
		Format: format,
	}

	return &i, nil
}
