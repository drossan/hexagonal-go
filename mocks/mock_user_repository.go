package mocks

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/drossan/core-api/domain/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"time"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id uint) (*model.User, error) {
	args := m.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetAll() ([]*model.User, error) {
	args := m.Called()
	return args.Get(0).([]*model.User), args.Error(1)
}

func (m *MockUserRepository) Paginate(page int, pageSize int) ([]*model.User, int, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]*model.User), args.Int(1), args.Error(2)
}

func (m *MockUserRepository) Delete(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Login(email, password string) (string, error) {
	user, err := m.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	ps := sha256.Sum256([]byte(password))
	pwd := fmt.Sprintf("%x", ps)

	if pwd == user.Password {
		claims := &jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
			"iss":     "Intranet API - generate by IslaIT",
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte("test_secret"))
	}

	return "", errors.New("invalid email or password")
}
