package mocks

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/stretchr/testify/mock"
)

type MockLevelPrivilegesRepository struct {
	mock.Mock
}

func (m *MockLevelPrivilegesRepository) CreateOrUpdate(levelPrivilege *model.LevelPrivileges) error {
	args := m.Called(levelPrivilege)
	return args.Error(0)
}

func (m *MockLevelPrivilegesRepository) GetAll() ([]*model.LevelPrivileges, error) {
	args := m.Called()
	return args.Get(0).([]*model.LevelPrivileges), args.Error(1)
}

func (m *MockLevelPrivilegesRepository) Delete(levelPrivilege *model.LevelPrivileges) error {
	args := m.Called(levelPrivilege)
	return args.Error(0)
}
