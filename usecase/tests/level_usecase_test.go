package usecase_test

import (
	"github.com/drossan/core-api/mocks"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/stretchr/testify/assert"
)

func TestLevelUseCase_CreateOrUpdateLevel(t *testing.T) {
	mockRepo := new(mocks.MockLevelRepository)
	mockLevel := &model.Level{Level: "Test Level"}

	mockRepo.On("CreateOrUpdate", mockLevel).Return(nil)

	uc := usecase.NewLevelUseCase(mockRepo)

	err := uc.CreateOrUpdateLevel(mockLevel)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestLevelUseCase_GetLevelByID(t *testing.T) {
	mockRepo := new(mocks.MockLevelRepository)
	mockLevel := &model.Level{Level: "Test Level"}

	mockRepo.On("GetByID", uint(1)).Return(mockLevel, nil)

	uc := usecase.NewLevelUseCase(mockRepo)

	level, err := uc.GetLevelByID(1)

	assert.Nil(t, err)
	assert.Equal(t, mockLevel, level)
	mockRepo.AssertExpectations(t)
}

func TestLevelUseCase_GetAllLevels(t *testing.T) {
	mockRepo := new(mocks.MockLevelRepository)
	mockLevels := []*model.Level{
		{Level: "Level1"},
		{Level: "Level2"},
	}

	mockRepo.On("GetAll").Return(mockLevels, nil)

	uc := usecase.NewLevelUseCase(mockRepo)

	levels, err := uc.GetAllLevels()

	assert.Nil(t, err)
	assert.Equal(t, mockLevels, levels)
	mockRepo.AssertExpectations(t)
}

func TestLevelUseCase_PaginateLevels(t *testing.T) {
	mockRepo := new(mocks.MockLevelRepository)
	mockLevels := []*model.Level{
		{Level: "Level1"},
		{Level: "Level2"},
	}

	mockRepo.On("Paginate", 1, 10).Return(mockLevels, 2, nil)

	uc := usecase.NewLevelUseCase(mockRepo)

	levels, total, err := uc.PaginateLevels(1, 10)

	assert.Nil(t, err)
	assert.Equal(t, mockLevels, levels)
	assert.Equal(t, 2, total)
	mockRepo.AssertExpectations(t)
}

func TestLevelUseCase_DeleteLevel(t *testing.T) {
	mockRepo := new(mocks.MockLevelRepository)
	mockLevel := &model.Level{Level: "Test Level"}

	mockRepo.On("Delete", mockLevel).Return(nil)

	uc := usecase.NewLevelUseCase(mockRepo)

	err := uc.DeleteLevel(mockLevel)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
