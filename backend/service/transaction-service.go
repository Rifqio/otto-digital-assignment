package service

import "voucher-app/dto"

// TransactionService is a struct to represent service of transaction
type TransactionService struct {
}

// NewTransactionService creates a new transaction service
func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

// CreateRedemptionTransaction is a function to create redemption transaction
func (t *TransactionService) CreateRedemptionTransaction(data dto.CreateTransactionRequest) error {
	return nil
}

// GetRedemptionTransactionDetail is a function to get redemption transaction detail
func (t *TransactionService) GetRedemptionTransactionDetail(id int) error {
	return nil
}
