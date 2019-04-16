package repository

import (
	"log"

	gpx "github.com/ptrv/go-gpx"
)

// NewGpxRepository initialize repository
func NewGpxRepository() *GpxRepository {
	data, err := gpx.ParseFile("./data/sample.gpx")

	if err != nil {
		log.Fatal(err)
	}

	var repository = GpxRepository{data}

	return &repository
}

// Repository interface defining what a repositiory should look like
type Repository interface {
	FindOne()
	FindAll()
}

// GpxRepository the GPX repository
type GpxRepository struct {
	data *gpx.Gpx
}

// FindOne find a single entry
func (r *GpxRepository) FindOne() *gpx.Gpx {
	return r.data
}

// FindAll find all entries
func (r *GpxRepository) FindAll() {

}
