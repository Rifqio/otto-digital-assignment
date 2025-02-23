package service

import (
	"testing"
	"time"
	"voucher-app/dto"
	"voucher-app/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVoucherRepository struct {
	mock.Mock
}

func (m *MockVoucherRepository) InsertVoucher(data dto.CreateVoucherRequest) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockVoucherRepository) FindVoucherByCode(code string) (*repositories.Voucher, error) {
	args := m.Called(code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repositories.Voucher), args.Error(1)
}

func (m *MockVoucherRepository) FindVoucherByBrand(brandID int) (*[]repositories.Voucher, error) {
	args := m.Called(brandID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*[]repositories.Voucher), args.Error(1)
}

func (m *MockVoucherRepository) FindVouchersByCodes(codes []string) ([]repositories.Voucher, error) {
	args := m.Called(codes)
	return args.Get(0).([]repositories.Voucher), args.Error(1)
}

func (m *MockVoucherRepository) FindVoucherByID(id int) (*repositories.Voucher, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repositories.Voucher), args.Error(1)
}

func TestCreateVoucher_Success(t *testing.T) {
	mockVoucherRepo := new(MockVoucherRepository)
	mockBrandRepo := new(MockBrandRepository)
	service := &VoucherService{
		voucherRepository: mockVoucherRepo,
		brandRepository:   mockBrandRepo,
	}

	validBrand := &repositories.Brand{
		ID:   1,
		Name: "Test Brand",
	}

	request := dto.CreateVoucherRequest{
		Code:      "TESTCODE",
		Point:     100,
		ValidAt:   time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
		BrandID:   1,
	}

	mockBrandRepo.On("FindBrandByID", 1).Return(validBrand, nil)
	mockVoucherRepo.On("FindVoucherByCode", "TESTCODE").Return(nil, nil)
	mockVoucherRepo.On("InsertVoucher", request).Return(nil)

	err := service.CreateVoucher(request)

	assert.NoError(t, err)
	mockBrandRepo.AssertExpectations(t)
	mockVoucherRepo.AssertExpectations(t)
}

func TestCreateVoucher_BrandNotFound(t *testing.T) {
	mockVoucherRepo := new(MockVoucherRepository)
	mockBrandRepo := new(MockBrandRepository)
	service := &VoucherService{
		voucherRepository: mockVoucherRepo,
		brandRepository:   mockBrandRepo,
	}

	request := dto.CreateVoucherRequest{
		Code:      "TESTCODE",
		Point:     100,
		ValidAt:   time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
		BrandID:   1,
	}

	mockBrandRepo.On("FindBrandByID", 1).Return(nil, nil)

	err := service.CreateVoucher(request)

	assert.EqualError(t, err, "brand not found")
	mockBrandRepo.AssertExpectations(t)
}

func TestCreateVoucher_VoucherExists(t *testing.T) {
	mockVoucherRepo := new(MockVoucherRepository)
	mockBrandRepo := new(MockBrandRepository)
	service := &VoucherService{
		voucherRepository: mockVoucherRepo,
		brandRepository:   mockBrandRepo,
	}

	validBrand := &repositories.Brand{
		ID:   1,
		Name: "Test Brand",
	}

	existingVoucher := &repositories.Voucher{
		Code: "TESTCODE",
	}

	request := dto.CreateVoucherRequest{
		Code:      "TESTCODE",
		Point:     100,
		ValidAt:   time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
		BrandID:   1,
	}

	mockBrandRepo.On("FindBrandByID", 1).Return(validBrand, nil)
	mockVoucherRepo.On("FindVoucherByCode", "TESTCODE").Return(existingVoucher, nil)

	err := service.CreateVoucher(request)

	assert.EqualError(t, err, "voucher already exists")
	mockBrandRepo.AssertExpectations(t)
	mockVoucherRepo.AssertExpectations(t)
}

func TestGetVoucher_Success(t *testing.T) {
	mockVoucherRepo := new(MockVoucherRepository)
	mockBrandRepo := new(MockBrandRepository)
	service := &VoucherService{
		voucherRepository: mockVoucherRepo,
		brandRepository:   mockBrandRepo,
	}

	expectedVoucher := &repositories.Voucher{
		ID:        1,
		Code:      "TESTCODE",
		Point:     100,
		ValidAt:   time.Now(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
		BrandID:   1,
	}

	mockVoucherRepo.On("FindVoucherByID", 1).Return(expectedVoucher, nil)

	voucher, err := service.GetVoucher(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedVoucher, voucher)
	mockVoucherRepo.AssertExpectations(t)
}

func TestGetVoucher_NotFound(t *testing.T) {
	mockVoucherRepo := new(MockVoucherRepository)
	mockBrandRepo := new(MockBrandRepository)
	service := &VoucherService{
		voucherRepository: mockVoucherRepo,
		brandRepository:   mockBrandRepo,
	}

	mockVoucherRepo.On("FindVoucherByID", 1).Return(nil, nil)

	voucher, err := service.GetVoucher(1)

	assert.EqualError(t, err, "voucher not found")
	assert.Nil(t, voucher)
	mockVoucherRepo.AssertExpectations(t)
}

func TestGetVoucherByBrand_Success(t *testing.T) {
	mockVoucherRepo := new(MockVoucherRepository)
	mockBrandRepo := new(MockBrandRepository)
	service := &VoucherService{
		voucherRepository: mockVoucherRepo,
		brandRepository:   mockBrandRepo,
	}

	expectedVouchers := &[]repositories.Voucher{
		{
			ID:        1,
			Code:      "TESTCODE1",
			Point:     100,
			ValidAt:   time.Now(),
			ExpiredAt: time.Now().Add(24 * time.Hour),
			BrandID:   1,
		},
		{
			ID:        2,
			Code:      "TESTCODE2",
			Point:     200,
			ValidAt:   time.Now(),
			ExpiredAt: time.Now().Add(24 * time.Hour),
			BrandID:   1,
		},
	}

	mockVoucherRepo.On("FindVoucherByBrand", 1).Return(expectedVouchers, nil)

	vouchers, err := service.GetVoucherByBrand(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedVouchers, vouchers)
	mockVoucherRepo.AssertExpectations(t)
}

func TestGetVoucherByBrand_NotFound(t *testing.T) {
	mockVoucherRepo := new(MockVoucherRepository)
	mockBrandRepo := new(MockBrandRepository)
	service := &VoucherService{
		voucherRepository: mockVoucherRepo,
		brandRepository:   mockBrandRepo,
	}

	mockVoucherRepo.On("FindVoucherByBrand", 1).Return(nil, nil)

	vouchers, err := service.GetVoucherByBrand(1)

	assert.EqualError(t, err, "voucher not found")
	assert.Nil(t, vouchers)
	mockVoucherRepo.AssertExpectations(t)
}
