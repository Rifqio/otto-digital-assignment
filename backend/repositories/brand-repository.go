package repositories

import (
	"fmt"
	"voucher-app/dto"
)

// BrandRepository is a struct to represent repository of brand
type BrandRepository struct {
}

// NewBrandRepository is a function to create new BrandRepository
func NewBrandRepository() BrandRepository {
	return BrandRepository{}
}

// InsertBrand is a function to insert brand
func (b BrandRepository) InsertBrand(data dto.CreateBrandRequest) error {
	fmt.Sprintf("Insert brand with name %s", data.BrandName)
	return nil
}
