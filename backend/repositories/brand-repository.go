package repositories

import (
	"time"
	"voucher-app/dto"

	"gorm.io/gorm"
)

// Brand struct is a model that represent brand table
type Brand struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"unique;not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// BrandRepository is a interface to represent repository of brand
type BrandRepository interface {
	InsertBrand(data dto.CreateBrandRequest) error
	FindBrandByName(name string) (*Brand, error)
	FindBrandByID(id int) (*Brand, error)
}

type brandRepositoryImpl struct {
	db *gorm.DB
}

// NewBrandRepository is a function to create new BrandRepository
func NewBrandRepository(db *gorm.DB) BrandRepository {
	return &brandRepositoryImpl{db: db}
}

// InsertBrand is a function to insert brand
func (b *brandRepositoryImpl) InsertBrand(data dto.CreateBrandRequest) error {
	brand := Brand{
		Name:      data.BrandName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result := b.db.Create(&brand)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// FindBrandByName is a function to find brand by name
func (b *brandRepositoryImpl) FindBrandByName(name string) (*Brand, error) {
	var brand Brand
	err := b.db.Where("name = ?", name).First(&brand).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &brand, nil
}

// FindBrandByID is a function to find brand by id
func (b *brandRepositoryImpl) FindBrandByID(id int) (*Brand, error) {
	var brand Brand
	err := b.db.Where("id = ?", id).First(&brand).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &brand, nil
}
