package repositories

// BrandRepository is a struct to represent repository of brand
type BrandRepository struct {
}

// NewBrandRepository is a function to create new BrandRepository
func NewBrandRepository() BrandRepository {
	return BrandRepository{}
}

// InsertBrand is a function to insert brand
func (b BrandRepository) InsertBrand() error {
	return nil
}
