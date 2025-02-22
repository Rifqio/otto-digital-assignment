package handler

import (
	"voucher-app/service"

	"github.com/labstack/echo/v4"
)

// TransactionHandler is a struct to represent handler of transaction
type TransactionHandler struct {
	transactionService *service.TransactionService
}

// NewTransactionHandler creates a new transaction handler
func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

// CreateRedemptionTransaction is a function to create redemption transaction
func (t *TransactionHandler) CreateRedemptionTransaction(e echo.Context) error {
	return nil
}

// GetRedemptionTransactionDetail is a function to get redemption transaction detail
func (t *TransactionHandler) GetRedemptionTransactionDetail(e echo.Context) error {
	return nil
}
