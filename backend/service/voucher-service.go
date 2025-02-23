package service

import (
	"errors"
	"voucher-app/dto"
	"voucher-app/repositories"

	"gorm.io/gorm"
)

// VoucherService is a struct to represent service of voucher
type VoucherService struct {
	voucherRepository repositories.VoucherRepository
	brandRepository   repositories.BrandRepository
}

// NewVoucherService is a function to create new VoucherService
func NewVoucherService(db *gorm.DB) *VoucherService {
	return &VoucherService{
		voucherRepository: repositories.NewVoucherRepository(db),
		brandRepository:   repositories.NewBrandRepository(db),
	}
}

// CreateVoucher is a function to create voucher
func (v VoucherService) CreateVoucher(data dto.CreateVoucherRequest) error {
	brand, err := v.brandRepository.FindBrandByID(data.BrandID)
	if err != nil {
		return err
	}

	if brand == nil {
		return errors.New("brand not found")
	}

	existingVoucher, err := v.voucherRepository.FindVoucherByCode(data.Code)
	if err != nil {
		return err
	}

	if existingVoucher != nil {
		return errors.New("voucher already exists")
	}

	err = v.voucherRepository.InsertVoucher(data)
	return nil
}

// GetVoucher is a function to get voucher
func (v VoucherService) GetVoucher(id int) (*repositories.Voucher, error) {
	voucher, err := v.voucherRepository.FindVoucherByID(id)
	if err != nil {
		return nil, err
	}

	if voucher == nil {
		return nil, errors.New("voucher not found")
	}

	return voucher, nil
}

// GetVoucherByBrand is a function to get voucher by brand
func (v VoucherService) GetVoucherByBrand(brandID int) (*[]repositories.Voucher, error) {
	voucher, err := v.voucherRepository.FindVoucherByBrand(brandID)
	if err != nil {
		return nil, err
	}

	if voucher == nil {
		return nil, errors.New("voucher not found")
	}

	return voucher, nil
}
