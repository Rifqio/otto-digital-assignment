package repositories

// UserRepository is a struct to represent repository of brand
type UserRepository struct {
}

// NewUserRepository is a function to create new UserRepository
func NewUserRepository() UserRepository {
	return UserRepository{}
}

// InsertVoucher is a function to insert brand
func (b UserRepository) InsertUser() error {
	return nil
}

func (b UserRepository) FindUserByID(id int) error {
	return nil
}
