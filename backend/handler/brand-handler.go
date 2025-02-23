package handler

import (
	"voucher-app/dto"
	"voucher-app/service"
	"voucher-app/utils"

	"github.com/labstack/echo/v4"
)

// BrandHandler is a struct to represent handler of brand
type BrandHandler struct {
	brandService *service.BrandService
}

// NewBrandHandler is a function to create new BrandHandler
func NewBrandHandler(brandService *service.BrandService) *BrandHandler {
	return &BrandHandler{brandService}
}

// CreateBrand is a function to create brand
func (b *BrandHandler) CreateBrand(c echo.Context) error {
	var req dto.CreateBrandRequest
	if err := c.Bind(&req); err != nil {
		return utils.ErrorResponse(c, 400, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return utils.ErrorResponse(c, 400, err.Error())
	}

	err := b.brandService.CreateBrand(req)
	if err != nil {
		if err.Error() == "brand name cannot be empty" || err.Error() == "brand already exists" {
			return utils.ErrorResponse(c, 400, err.Error())
		}
		return utils.ErrorResponse(c, 500, "failed to create brand")
	}

	return utils.SuccessResponse(c, "Brand created", nil)
}
