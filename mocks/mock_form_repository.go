package mocks

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockFormRepository struct {
	mock.Mock
}

func (m *MockFormRepository) CreateOrUpdate(form *model.Form) error {
	args := m.Called(form)
	return args.Error(0)
}

func (m *MockFormRepository) GetByID(id uint) (*model.Form, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Form), args.Error(1)
}

func (m *MockFormRepository) GetAll() ([]*model.Form, error) {
	args := m.Called()
	return args.Get(0).([]*model.Form), args.Error(1)
}

func (m *MockFormRepository) Paginate(page, pageSize int) ([]*model.Form, int, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]*model.Form), args.Int(1), args.Error(2)
}

func (m *MockFormRepository) Delete(form *model.Form) error {
	args := m.Called(form)
	return args.Error(0)
}
