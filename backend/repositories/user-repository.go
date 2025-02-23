package repositories

import (
	"time"

	"gorm.io/gorm"
)

// User represents the model for users
type User struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoincrement"`
	Email     string    `json:"email" gorm:"unique;not null"`
	FullName  string    `json:"fullName" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// UserRepository defines the interface for user repository operations
type UserRepository interface {
	FindUserByEmail(email string) (*User, error)
}

// userRepositoryImpl is the concrete implementation of UserRepository
type userRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// FindUserByEmail finds a user by email
func (u *userRepositoryImpl) FindUserByEmail(email string) (*User, error) {
	var user User
	err := u.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
