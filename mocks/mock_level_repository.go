package mocks

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockLevelRepository struct {
	mock.Mock
}

func (m *MockLevelRepository) CreateOrUpdate(level *model.Level) error {
	args := m.Called(level)
	return args.Error(0)
}

func (m *MockLevelRepository) GetByID(id uint) (*model.Level, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Level), args.Error(1)
}

func (m *MockLevelRepository) GetAll() ([]*model.Level, error) {
	args := m.Called()
	return args.Get(0).([]*model.Level), args.Error(1)
}

func (m *MockLevelRepository) Paginate(page, pageSize int) ([]*model.Level, int, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]*model.Level), args.Int(1), args.Error(2)
}

func (m *MockLevelRepository) Delete(level *model.Level) error {
	args := m.Called(level)
	return args.Error(0)
}
