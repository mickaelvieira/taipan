package repository

// NewGpxRepository initialize repository
func NewGpxRepository() *GpxRepository {
	// data, err := gpx.ParseFile("./data/sample.gpx")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	var repository = GpxRepository{}

	return &repository
}

// Repository interface defining what a repositiory should look like
type Repository interface {
	FindOne()
	FindAll()
}

// GpxRepository the GPX repository
type GpxRepository struct{}

// FindOne find a single entry
func (r *GpxRepository) FindOne() {
}

// FindAll find all entries
func (r *GpxRepository) FindAll() {
}
