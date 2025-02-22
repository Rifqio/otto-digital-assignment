package dto

type CreateBrandRequest struct {
	BrandName string `json:"brandName" validate:"required"`
}
