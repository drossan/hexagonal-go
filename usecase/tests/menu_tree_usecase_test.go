package usecase_test

import (
	"github.com/drossan/core-api/mocks"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/stretchr/testify/assert"
)

func TestMenuTreeUseCase_CreateOrUpdateMenuTree(t *testing.T) {
	mockRepo := new(mocks.MockMenuTreeRepository)
	mockMenu := &model.MenuTree{Title: "Test Menu"}

	mockRepo.On("CreateOrUpdate", mockMenu).Return(nil)

	uc := usecase.NewMenuTreeUseCase(mockRepo)

	err := uc.CreateOrUpdateMenuTree(mockMenu)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMenuTreeUseCase_GetAllMenuTrees(t *testing.T) {
	mockRepo := new(mocks.MockMenuTreeRepository)
	mockMenus := []*model.MenuTree{
		{Title: "Menu1"},
		{Title: "Menu2"},
	}

	mockRepo.On("GetAll").Return(mockMenus, nil)

	uc := usecase.NewMenuTreeUseCase(mockRepo)

	menus, err := uc.GetAllMenuTrees()

	assert.Nil(t, err)
	assert.Equal(t, mockMenus, menus)
	mockRepo.AssertExpectations(t)
}

func TestMenuTreeUseCase_DeleteMenuTree(t *testing.T) {
	mockRepo := new(mocks.MockMenuTreeRepository)
	mockMenu := &model.MenuTree{Title: "Test Menu"}

	mockRepo.On("Delete", mockMenu).Return(nil)

	uc := usecase.NewMenuTreeUseCase(mockRepo)

	err := uc.DeleteMenuTree(mockMenu)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
