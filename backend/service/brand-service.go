package service

import (
	"voucher-app/dto"
	"voucher-app/repositories"
)

// BrandService is a struct to represent service of brand
type BrandService struct {
	brandRepository repositories.BrandRepository
}

// NewBrandService is a function to create new BrandService
func NewBrandService() *BrandService {
	return &BrandService{
		brandRepository: repositories.NewBrandRepository(),
	}
}

// CreateBrand is a function to create brand
func (b BrandService) CreateBrand(data dto.CreateBrandRequest) error {
	err := b.brandRepository.InsertBrand()
	if err != nil {
		return err
	}
	return nil
}
