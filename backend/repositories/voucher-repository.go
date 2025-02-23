package repositories

import (
	"time"
	"voucher-app/dto"

	"gorm.io/gorm"
)

// Voucher is a struct to represent the model
type Voucher struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoincrement"`
	Code      string    `json:"code" gorm:"unique;not null"`
	Point     int       `json:"point" gorm:"not null"`
	ValidAt   time.Time `json:"validAt" gorm:"not null"`
	ExpiredAt time.Time `json:"expiredAt" gorm:"not null"`
	BrandID   int       `json:"brandId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// VoucherRepository is a struct to represent repository of brand
type VoucherRepository struct {
	db *gorm.DB
}

// NewVoucherRepository is a function to create new VoucherRepository
func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return VoucherRepository{
		db: db,
	}
}

// InsertVoucher is a function to insert brand
func (b VoucherRepository) InsertVoucher(data dto.CreateVoucherRequest) error {
	voucher := Voucher{
		Code:      data.Code,
		Point:     data.Point,
		ValidAt:   data.ValidAt,
		ExpiredAt: data.ExpiredAt,
		BrandID:   data.BrandID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := b.db.Create(&voucher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindVoucherByCode is a function to find voucher by code
func (b VoucherRepository) FindVoucherByCode(code string) (*Voucher, error) {
	var voucher Voucher
	result := b.db.Where("code = ?", code).First(&voucher)
	if result.Error != nil {
		return nil, result.Error
	}
	return &voucher, nil
}

// FindVoucherByBrand is a function to find voucher by brand
func (b VoucherRepository) FindVoucherByBrand(brandID int) (*[]Voucher, error) {
	var vouchers []Voucher
	result := b.db.Where("brand_id = ?", brandID).Find(&vouchers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &vouchers, nil
}

// FindVouchersByCodes retrieves multiple vouchers based on a slice of voucher codes.
func (b VoucherRepository) FindVouchersByCodes(codes []string) ([]Voucher, error) {
	var vouchers []Voucher
	result := b.db.Where("code IN (?)", codes).Find(&vouchers)
	if result.Error != nil {
		return nil, result.Error
	}
	return vouchers, nil
}

// FindVoucherByID is a function to find voucher by ID
func (b VoucherRepository) FindVoucherByID(id int) (*Voucher, error) {
	var voucher Voucher
	result := b.db.Where("id = ?", id).First(&voucher)
	if result.Error != nil {
		return nil, result.Error
	}
	return &voucher, nil
}
