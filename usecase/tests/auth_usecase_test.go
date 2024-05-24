package usecase_test

import (
	"crypto/sha256"
	"fmt"
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/mocks"
	"github.com/drossan/core-api/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserUseCase_Login(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	password := "password"
	ps := sha256.Sum256([]byte(password))
	pwd := fmt.Sprintf("%x", ps)

	mockUser := &model.User{
		Email:    "test@example.com",
		Password: pwd,
	}

	mockRepo.On("GetByEmail", "test@example.com").Return(mockUser, nil)

	uc := usecase.NewUserUseCase(mockRepo)

	token, err := uc.Login("test@example.com", password)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestAuthUseCase_LoginInvalidPassword(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	password := "password"
	ps := sha256.Sum256([]byte(password))
	pwd := fmt.Sprintf("%x", ps)

	mockUser := &model.User{
		Email:    "test@example.com",
		Password: pwd,
	}

	mockRepo.On("GetByEmail", "test@example.com").Return(mockUser, nil)

	uc := usecase.NewUserUseCase(mockRepo)

	token, err := uc.Login("test@example.com", "wrongpassword")

	assert.NotNil(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}
