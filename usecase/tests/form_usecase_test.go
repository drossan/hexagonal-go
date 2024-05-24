package usecase_test

import (
	"github.com/drossan/core-api/mocks"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/usecase"
	"github.com/stretchr/testify/assert"
)

func TestFormUseCase_CreateOrUpdateForm(t *testing.T) {
	mockRepo := new(mocks.MockFormRepository)
	mockForm := &model.Form{Title: "Test Form"}

	mockRepo.On("CreateOrUpdate", mockForm).Return(nil)

	uc := usecase.NewFormUseCase(mockRepo)

	err := uc.CreateOrUpdateForm(mockForm)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFormUseCase_GetAllForms(t *testing.T) {
	mockRepo := new(mocks.MockFormRepository)
	mockForms := []*model.Form{
		{Title: "Form1"},
		{Title: "Form2"},
	}

	mockRepo.On("GetAll").Return(mockForms, nil)

	uc := usecase.NewFormUseCase(mockRepo)

	forms, err := uc.GetAllForms()

	assert.Nil(t, err)
	assert.Equal(t, mockForms, forms)
	mockRepo.AssertExpectations(t)
}

func TestFormUseCase_DeleteForm(t *testing.T) {
	mockRepo := new(mocks.MockFormRepository)
	mockForm := &model.Form{Title: "Test Form"}

	mockRepo.On("Delete", mockForm).Return(nil)

	uc := usecase.NewFormUseCase(mockRepo)

	err := uc.DeleteForm(mockForm)

	assert.Nil(t, err)
	mockRepo.AssertExpectations(t)
}
