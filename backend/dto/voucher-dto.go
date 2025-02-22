package dto

import "time"

// CreateVoucherRequest is a struct to represent request payload to create voucher
type CreateVoucherRequest struct {
	Code      string     `json:"code"`
	Point     int        `json:"point"`
	ValidAt   time.Time  `json:"validAt"`
	ExpiredAt time.Time  `json:"expiredAt"`
	BrandID   int        `json:"brandId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
