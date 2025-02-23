package repositories

import (
	"time"
)

// Voucher is a struct to represent the model
type Voucher struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoincrement"`
	Code      string    `json:"code" gorm:"unique;not null"`
	Point     int       `json:"point" gorm:"not null"`
	ValidAt   time.Time `json:"validAt gorm:"not null"`
	ExpiredAt time.Time
	BrandID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// VoucherRepository is a struct to represent repository of brand
type VoucherRepository struct {
}

// NewVoucherRepository is a function to create new VoucherRepository
func NewVoucherRepository() VoucherRepository {
	return VoucherRepository{}
}

// InsertVoucher is a function to insert brand
func (b VoucherRepository) InsertVoucher() error {
	return nil
}
