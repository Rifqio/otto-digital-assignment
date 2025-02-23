package service

import (
	"errors"
	"fmt"
	"time"
	"voucher-app/dto"
	"voucher-app/repositories"

	"gorm.io/gorm"
)

// TransactionService is a struct to represent service of transaction
type TransactionService struct {
	transactionRepository repositories.TransactionRepository
	userRepository        repositories.UserRepository
	voucherRepository     repositories.VoucherRepository
}

// NewTransactionService creates a new transaction service
func NewTransactionService(db *gorm.DB) *TransactionService {
	return &TransactionService{
		transactionRepository: repositories.NewTransactionRepository(db),
		userRepository:        repositories.NewUserRepository(db),
		voucherRepository:     repositories.NewVoucherRepository(db),
	}
}

// CreateRedemptionTransaction handles voucher redemption and logs transaction history
func (t *TransactionService) CreateRedemptionTransaction(data dto.CreateTransactionRequest) error {
	// Validate customer email
	customer, err := t.userRepository.FindUserByEmail(data.CustomerEmail)
	if err != nil {
		return err
	}

	if customer == nil {
		return errors.New("customer not found")
	}

	// Validate voucher codes
	var voucherCodes []string
	voucherQuantities := make(map[string]int)

	for _, v := range data.Vouchers {
		voucherCodes = append(voucherCodes, v.VoucherCode)
		voucherQuantities[v.VoucherCode] = v.Quantity
	}

	vouchers, err := t.voucherRepository.FindVouchersByCodes(voucherCodes)
	if err != nil {
		return err
	}

	if len(vouchers) != len(voucherCodes) {
		return errors.New("one or more vouchers are invalid")
	}

	// Calculate total points and insert into transaction history
	totalPoints := 0
	currentTime := time.Now()

	for _, voucher := range vouchers {
		if currentTime.Before(voucher.ValidAt) || currentTime.After(voucher.ExpiredAt) {
			return fmt.Errorf("voucher %s is not valid", voucher.Code)
		}

		quantity := voucherQuantities[voucher.Code]
		pointsRedeemed := voucher.Point * quantity
		totalPoints += pointsRedeemed

		// TRX-<customer_id>-<voucher_id>-<timestamp>
		trxID := fmt.Sprintf("TRX-%d-%d-%d", customer.ID, voucher.ID, currentTime.Unix())
		transactionHistory := repositories.TransactionHistory{
			ID:             trxID,
			VoucherID:      voucher.ID,
			UserID:         customer.ID,
			PointsRedeemed: pointsRedeemed,
			CreatedAt:      int(time.Now().Unix()),
			UpdatedAt:      int(time.Now().Unix()),
		}

		err := t.transactionRepository.CreateTransactionHistory(transactionHistory)
		if err != nil {
			return fmt.Errorf("failed to save transaction history: %w", err)
		}
	}

	return nil
}

// GetRedemptionTransactionDetail is a function to get redemption transaction detail
func (t *TransactionService) GetRedemptionTransactionDetail(id string) (*repositories.TransactionHistory, error) {
	transaction, err := t.transactionRepository.FindTransactionHistoryByID(id)
	if err != nil {
		return nil, err
	}

	if transaction == nil {
		return nil, errors.New("transaction not found")
	}

	return transaction, nil
}
