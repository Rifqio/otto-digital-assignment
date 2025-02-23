package repositories

import (
	"time"

	"gorm.io/gorm"
)

// User is a struct to represent the model
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoincrement"`
	Email     string    `json:"email" gorm:"unique;notnull"`
	FullName  string    `json:"fullName" gorm:"notnull"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UserRepository is a struct to represent repository of brand
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is a function to create new UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

// InsertUser is a function to insert user
func (b UserRepository) InsertUser() error {
	return nil
}

// FindUserByEmail is a function to find user by email
func (b UserRepository) FindUserByEmail(email string) (*User, error) {
	var user User
	err := b.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
