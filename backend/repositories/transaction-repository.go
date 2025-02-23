package repositories

import "gorm.io/gorm"

// TransactionHistory represents the model for transaction history
type TransactionHistory struct {
	ID             string `json:"id" gorm:"primaryKey"`
	VoucherID      uint   `json:"voucherId" gorm:"not null"`
	UserID         uint   `json:"userId" gorm:"not null"`
	PointsRedeemed int    `json:"pointsRedeemed" gorm:"not null"`
	CreatedAt      int    `json:"createdAt" gorm:"not null"`
	UpdatedAt      int    `json:"updatedAt" gorm:"not null"`
}

// TransactionRepository defines the methods that the repository must implement
type TransactionRepository interface {
	CreateTransactionHistory(data TransactionHistory) error
	FindTransactionHistoryByID(id string) (*TransactionHistory, error)
}

// transactionRepositoryImpl is the concrete implementation of TransactionRepository
type transactionRepositoryImpl struct {
	db *gorm.DB
}

// NewTransactionRepository creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepositoryImpl{db: db}
}

// CreateTransactionHistory inserts a record into the transaction_history table
func (t *transactionRepositoryImpl) CreateTransactionHistory(data TransactionHistory) error {
	result := t.db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindTransactionHistoryByID finds a transaction history by its ID
func (t *transactionRepositoryImpl) FindTransactionHistoryByID(id string) (*TransactionHistory, error) {
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
