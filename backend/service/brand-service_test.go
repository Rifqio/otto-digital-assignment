package service

import (
	"errors"
	"testing"
	"voucher-app/dto"
	"voucher-app/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBrandRepository is a mock implementation of BrandRepository
type MockBrandRepository struct {
	mock.Mock
}

func (m *MockBrandRepository) InsertBrand(data dto.CreateBrandRequest) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockBrandRepository) FindBrandByName(name string) (*repositories.Brand, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repositories.Brand), args.Error(1)
}

func (m *MockBrandRepository) FindBrandByID(id int) (*repositories.Brand, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*repositories.Brand), args.Error(1)
}

func TestCreateBrand_Success(t *testing.T) {
	mockRepo := new(MockBrandRepository)
	service := &BrandService{brandRepository: mockRepo}

	request := dto.CreateBrandRequest{BrandName: "Nike"}
	mockRepo.On("FindBrandByName", "Nike").Return(nil, nil)
	mockRepo.On("InsertBrand", request).Return(nil)

	err := service.CreateBrand(request)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateBrand_EmptyName(t *testing.T) {
	mockRepo := new(MockBrandRepository)
	service := &BrandService{brandRepository: mockRepo}

	request := dto.CreateBrandRequest{BrandName: " "}
	err := service.CreateBrand(request)

	assert.EqualError(t, err, "brand name cannot be empty")
}

func TestCreateBrand_AlreadyExists(t *testing.T) {
	mockRepo := new(MockBrandRepository)
	service := &BrandService{brandRepository: mockRepo}

	existingBrand := &repositories.Brand{Name: "Nike"}
	request := dto.CreateBrandRequest{BrandName: "Nike"}

	mockRepo.On("FindBrandByName", "Nike").Return(existingBrand, nil)
	err := service.CreateBrand(request)

	assert.EqualError(t, err, "brand already exists")
	mockRepo.AssertExpectations(t)
}

func TestCreateBrand_RepositoryError(t *testing.T) {
	mockRepo := new(MockBrandRepository)
	service := &BrandService{brandRepository: mockRepo}

	request := dto.CreateBrandRequest{BrandName: "Nike"}
	mockRepo.On("FindBrandByName", "Nike").Return(nil, errors.New("database error"))

	err := service.CreateBrand(request)

	assert.EqualError(t, err, "database error")
	mockRepo.AssertExpectations(t)
}
