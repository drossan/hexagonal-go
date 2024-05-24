package mocks

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockMenuTreeRepository struct {
	mock.Mock
}

func (m *MockMenuTreeRepository) CreateOrUpdate(MenuTree *model.MenuTree) error {
	args := m.Called(MenuTree)
	return args.Error(0)
}

func (m *MockMenuTreeRepository) GetByID(id uint) (*model.MenuTree, error) {
	args := m.Called(id)
	return args.Get(0).(*model.MenuTree), args.Error(1)
}

func (m *MockMenuTreeRepository) GetAll() ([]*model.MenuTree, error) {
	args := m.Called()
	return args.Get(0).([]*model.MenuTree), args.Error(1)
}

func (m *MockMenuTreeRepository) Paginate(page, pageSize int) ([]*model.MenuTree, int, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]*model.MenuTree), args.Int(1), args.Error(2)
}

func (m *MockMenuTreeRepository) Delete(MenuTree *model.MenuTree) error {
	args := m.Called(MenuTree)
	return args.Error(0)
}
