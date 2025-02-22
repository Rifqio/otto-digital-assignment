package service

import "voucher-app/repositories"

// VoucherService is a struct to represent service of voucher
type VoucherService struct {
	voucherRepository repositories.VoucherRepository
}

// NewVoucherService is a function to create new VoucherService
func NewVoucherService() *VoucherService {
	return &VoucherService{
		voucherRepository: repositories.NewVoucherRepository(),
	}
}

// CreateVoucher is a function to create voucher
func (v VoucherService) CreateVoucher() error {
	return nil
}

// GetVoucher is a function to get voucher
func (v VoucherService) GetVoucher(id int) error {
	return nil
}

// GetVoucherByBrand is a function to get voucher by brand
func (v VoucherService) GetVoucherByBrand(brandID int) error {
	return nil
}
