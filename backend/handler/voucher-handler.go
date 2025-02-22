package handler

import (
	"strconv"
	"voucher-app/dto"
	"voucher-app/service"
	"voucher-app/utils"

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
	var req dto.CreateVoucherRequest
	if err := e.Bind(&req); err != nil {
		return utils.ErrorResponse(e, 400, err.Error())
	}

	err := v.voucherService.CreateVoucher()
	if err != nil {
		return utils.ErrorResponse(e, 500, err.Error())
	}
	return utils.CreatedResponse(e, "Voucher created")
}

// GetVoucher is a function to get voucher
func (v *VoucherHandler) GetVoucher(e echo.Context) error {
	voucherID := e.QueryParam("id")
	if voucherID == "" {
		return utils.ErrorResponse(e, 400, "Voucher ID is required")
	}

	parsedVoucherID, err := strconv.Atoi(voucherID)
	if err != nil {
		return utils.ErrorResponse(e, 400, "Voucher ID must be a number")
	}

	err = v.voucherService.GetVoucher(parsedVoucherID)
	return nil
}

// GetVoucherByBrand is a function to get voucher by brand
func (v *VoucherHandler) GetVoucherByBrand(e echo.Context) error {
	brandID := e.QueryParam("id")
	if brandID == "" {
		return utils.ErrorResponse(e, 400, "Brand ID is required")
	}

	parsedBrandID, err := strconv.Atoi(brandID)
	if err != nil {
		return utils.ErrorResponse(e, 400, "Brand ID must be a number")
	}

	err = v.voucherService.GetVoucherByBrand(parsedBrandID)
	return nil
}
