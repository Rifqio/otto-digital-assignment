package repositories

import (
	"time"
	"voucher-app/dto"

	"gorm.io/gorm"
)

// Voucher represents the voucher model
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

// VoucherRepository defines the interface for voucher repository operations
type VoucherRepository interface {
	InsertVoucher(data dto.CreateVoucherRequest) error
	FindVoucherByCode(code string) (*Voucher, error)
	FindVoucherByBrand(brandID int) (*[]Voucher, error)
	FindVouchersByCodes(codes []string) ([]Voucher, error)
	FindVoucherByID(id int) (*Voucher, error)
}

// voucherRepositoryImpl is the concrete implementation of VoucherRepository
type voucherRepositoryImpl struct {
	db *gorm.DB
}

// NewVoucherRepository creates a new instance of VoucherRepository
func NewVoucherRepository(db *gorm.DB) VoucherRepository {
	return &voucherRepositoryImpl{db: db}
}

// InsertVoucher inserts a new voucher record into the database
func (v *voucherRepositoryImpl) InsertVoucher(data dto.CreateVoucherRequest) error {
	voucher := Voucher{
		Code:      data.Code,
		Point:     data.Point,
		ValidAt:   data.ValidAt,
		ExpiredAt: data.ExpiredAt,
		BrandID:   data.BrandID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := v.db.Create(&voucher)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindVoucherByCode retrieves a voucher by its code
func (v *voucherRepositoryImpl) FindVoucherByCode(code string) (*Voucher, error) {
	var voucher Voucher
	result := v.db.Where("code = ?", code).First(&voucher)
	if result.Error != nil {
		return nil, result.Error
	}
	return &voucher, nil
}

// FindVoucherByBrand retrieves vouchers associated with a specific brand
func (v *voucherRepositoryImpl) FindVoucherByBrand(brandID int) (*[]Voucher, error) {
	var vouchers []Voucher
	result := v.db.Where("brand_id = ?", brandID).Find(&vouchers)
	if result.Error != nil {
		return nil, result.Error
	}
	return &vouchers, nil
}

// FindVouchersByCodes retrieves multiple vouchers based on a list of voucher codes
func (v *voucherRepositoryImpl) FindVouchersByCodes(codes []string) ([]Voucher, error) {
	var vouchers []Voucher
	result := v.db.Where("code IN (?)", codes).Find(&vouchers)
	if result.Error != nil {
		return nil, result.Error
	}
	return vouchers, nil
}

// FindVoucherByID retrieves a voucher by its ID
func (v *voucherRepositoryImpl) FindVoucherByID(id int) (*Voucher, error) {
	var voucher Voucher
	result := v.db.Where("id = ?", id).First(&voucher)
	if result.Error != nil {
		return nil, result.Error
	}
	return &voucher, nil
}
