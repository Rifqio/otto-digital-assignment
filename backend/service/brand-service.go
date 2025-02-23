package service

import (
	"errors"
	"strings"
	"voucher-app/dto"
	"voucher-app/repositories"

	"gorm.io/gorm"
)

// BrandService is a struct to represent service of brand
type BrandService struct {
	brandRepository repositories.BrandRepository
}

// NewBrandService is a function to create new BrandService
func NewBrandService(db *gorm.DB) *BrandService {
	return &BrandService{
		brandRepository: repositories.NewBrandRepository(db),
	}
}

// CreateBrand is a function to create brand
func (b *BrandService) CreateBrand(data dto.CreateBrandRequest) error {
	data.BrandName = strings.TrimSpace(data.BrandName)

	if data.BrandName == "" {
		return errors.New("brand name cannot be empty")
	}

	existingBrand, err := b.brandRepository.FindBrandByName(data.BrandName)
	if err != nil {
		return err
	}

	if existingBrand != nil {
		return errors.New("brand already exists")
	}

	return b.brandRepository.InsertBrand(data)

}
