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

// BrandRepository is a struct to represent repository of brand
type BrandRepository struct {
	db *gorm.DB
}

// NewBrandRepository is a function to create new BrandRepository
func NewBrandRepository(db *gorm.DB) BrandRepository {
	return BrandRepository{
		db: db,
	}
}

// InsertBrand is a function to insert brand
func (b BrandRepository) InsertBrand(data dto.CreateBrandRequest) error {
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

func (b BrandRepository) FindBrandByName(name string) (*Brand, error) {
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
