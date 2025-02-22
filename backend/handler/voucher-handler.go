package handler

import (
	"voucher-app/service"

	"github.com/labstack/echo/v4"
)

// VoucherHandler is a struct to represent handler of voucher
type VoucherHandler struct {
	voucherService *service.VoucherService
}

// NewVoucherHandler is a function to create new VoucherHandler
func NewVoucherHandler(voucherService *service.VoucherService) *VoucherHandler {
	return &VoucherHandler{voucherService}
}

// CreateVoucher is a function to create voucher
func (v *VoucherHandler) CreateVoucher(e echo.Context) error {
	return nil
}

// GetVoucher is a function to get voucher
func (v *VoucherHandler) GetVoucher(e echo.Context) error {
	return nil
}

// GetVoucherByBrand is a function to get voucher by brand
func (v *VoucherHandler) GetVoucherByBrand(e echo.Context) error {
	return nil
}
