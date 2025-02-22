package repositories

// VoucherRepository is a struct to represent repository of brand
type VoucherRepository struct {
}

// NewVoucherRepository is a function to create new VoucherRepository
func NewVoucherRepository() VoucherRepository {
	return VoucherRepository{}
}

// InsertVoucher is a function to insert brand
func (b VoucherRepository) InsertVoucher() error {
	return nil
}
