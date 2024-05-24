package usecase_test

import (
	"github.com/drossan/core-api/mocks"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/stretchr/testify/assert"
)

func TestLevelPrivilegesUseCase_CreateOrUpdateLevelPrivilege(t *testing.T) {
	mockRepo := new(mocks.MockLevelPrivilegesRepository)
	mockLevelPrivilege := &model.LevelPrivileges{FormID: 1, Read: true, Write: true}

	mockRepo.On("CreateOrUpdate", mockLevelPrivilege).Return(nil)

	uc := usecase.NewLevelPrivilegesUseCase(mockRepo)

	err := uc.CreateOrUpdateLevelPrivilege(mockLevelPrivilege)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLevelPrivilegesUseCase_GetAllLevelPrivileges(t *testing.T) {
	mockRepo := new(mocks.MockLevelPrivilegesRepository)
	mockLevelPrivileges := []*model.LevelPrivileges{
		{FormID: 1, Read: true, Write: true},
		{FormID: 2, Read: true, Write: false},
	}

	mockRepo.On("GetAll").Return(mockLevelPrivileges, nil)

	uc := usecase.NewLevelPrivilegesUseCase(mockRepo)

	levelPrivileges, err := uc.GetAllLevelPrivilege()

	assert.Nil(t, err)
	assert.Equal(t, mockLevelPrivileges, levelPrivileges)
	mockRepo.AssertExpectations(t)
}

func TestLevelPrivilegesUseCase_DeleteLevelPrivilege(t *testing.T) {
	mockRepo := new(mocks.MockLevelPrivilegesRepository)
	mockLevelPrivilege := &model.LevelPrivileges{FormID: 1, Read: true, Write: true}

	mockRepo.On("Delete", mockLevelPrivilege).Return(nil)

	uc := usecase.NewLevelPrivilegesUseCase(mockRepo)

	err := uc.DeleteLevelPrivilege(mockLevelPrivilege)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
