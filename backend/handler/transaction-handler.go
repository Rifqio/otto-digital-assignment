package handler

import (
	"strconv"
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
		return utils.ErrorResponse(e, 400, err.Error())
	}

	err := t.transactionService.CreateRedemptionTransaction(req)
	if err != nil {
		return utils.ErrorResponse(e, 500, err.Error())
	}

	return utils.CreatedResponse(e, "Redemption transaction created")
}

// GetRedemptionTransactionDetail is a function to get redemption transaction detail
func (t *TransactionHandler) GetRedemptionTransactionDetail(e echo.Context) error {
	transactionID := e.QueryParam("transactionId")
	if transactionID == "" {
		return utils.ErrorResponse(e, 400, "Transaction ID is required")
	}

	parsedTransactionID, err := strconv.Atoi(transactionID)
	if err != nil {
		return utils.ErrorResponse(e, 400, "Transaction ID must be a number")
	}

	err = t.transactionService.GetRedemptionTransactionDetail(parsedTransactionID)
	if err != nil {
		return utils.ErrorResponse(e, 500, err.Error())
	}

	return utils.SuccessResponse(e, "Redemption transaction detail retrieved", nil)
}
