package dto

// VoucherRedeem is a struct to represent request payload to redeem voucher
type VoucherRedeem struct {
	VoucherCode string `json:"voucherCode"`
	Quantity    int    `json:"quantity"`
}

// CreateTransactionRequest is a struct to represent request payload to create transaction
type CreateTransactionRequest struct {
	CustomerEmail string          `json:"customerEmail"`
	Vouchers      []VoucherRedeem `json:"vouchers"`
}
