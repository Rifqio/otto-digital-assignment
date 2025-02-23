package handler

import (
	"strings"
	"voucher-app/dto"
	"voucher-app/service"
	"voucher-app/utils"

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
	var req dto.CreateTransactionRequest

	if err := e.Bind(&req); err != nil {
		return utils.ErrorResponse(e, 400, "invalid request payload")
	}

	if err := e.Validate(req); err != nil {
		return utils.ErrorResponse(e, 400, err.Error())
	}

	err := t.transactionService.CreateRedemptionTransaction(req)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return utils.ErrorResponse(e, 404, err.Error())
		} else if strings.Contains(err.Error(), "invalid") {
			return utils.ErrorResponse(e, 400, err.Error())
		}
		return utils.ErrorResponse(e, 500, "failed to process transaction")
	}

	return utils.CreatedResponse(e, "Redemption transaction created successfully")
}

// GetRedemptionTransactionDetail is a function to get redemption transaction detail
func (t *TransactionHandler) GetRedemptionTransactionDetail(e echo.Context) error {
	transactionID := e.QueryParam("transactionId")
	if transactionID == "" {
		return utils.ErrorResponse(e, 400, "Transaction ID is required")
	}

	transaction, err := t.transactionService.GetRedemptionTransactionDetail(transactionID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse(e, 404, "Transaction not found")
		}
		return utils.ErrorResponse(e, 500, err.Error())
	}

	return utils.SuccessResponse(e, "Redemption transaction detail retrieved", transaction)
}
