package repositories

import "gorm.io/gorm"

// TransactionHistory is a struct to represent the model
type TransactionHistory struct {
	ID             string `json:"id" gorm:"primaryKey"`
	VoucherID      uint   `json:"voucherId" gorm:"not null"`
	UserID         uint   `json:"userId" gorm:"not null"`
	PointsRedeemed int    `json:"pointsRedeemed" gorm:"not null"`
	CreatedAt      int    `json:"createdAt" gorm:"not null"`
	UpdatedAt      int    `json:"updatedAt" gorm:"not null"`
}

// TransactionRepository is a struct to represent repository of brand
type TransactionRepository struct {
	db *gorm.DB
}

// NewTransactionRepository is a function to create new TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return TransactionRepository{
		db: db,
	}
}

// CreateTransactionHistory inserts a record into the transaction_history table
func (t TransactionRepository) CreateTransactionHistory(data TransactionHistory) error {
	result := t.db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindTransactionHistoryByID finds a transaction history by its ID
func (t TransactionRepository) FindTransactionHistoryByID(id string) (*TransactionHistory, error) {
	var transactionHistory TransactionHistory
	err := t.db.Where("id = ?", id).First(&transactionHistory).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &transactionHistory, nil
}
