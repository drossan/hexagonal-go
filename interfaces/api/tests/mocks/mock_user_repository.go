package mocks

import (
	"errors"
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MockUserRepository struct {
	CreateFunc     func(user *model.User) error
	UpdateFunc     func(user *model.User) error
	GetByIDFunc    func(id uint) (*model.User, error)
	GetByEmailFunc func(email string) (*model.User, error)
	GetAllFunc     func() ([]*model.User, error)
	PaginateFunc   func(page int, pageSize int) ([]*model.User, int, error)
	DeleteFunc     func(user *model.User) error
	LoginFunc      func(email, password string) (string, error)
}

var _ repository.UserRepository = &MockUserRepository{}

func (m *MockUserRepository) Create(user *model.User) error {
	return m.CreateFunc(user)
}

func (m *MockUserRepository) Update(user *model.User) error {
	return m.UpdateFunc(user)
}

func (m *MockUserRepository) GetByID(id uint) (*model.User, error) {
	return m.GetByIDFunc(id)
}

func (m *MockUserRepository) GetByEmail(email string) (*model.User, error) {
	return m.GetByEmailFunc(email)
}

func (m *MockUserRepository) GetAll() ([]*model.User, error) {
	return m.GetAllFunc()
}

func (m *MockUserRepository) Paginate(page int, pageSize int) ([]*model.User, int, error) {
	return m.PaginateFunc(page, pageSize)
}

func (m *MockUserRepository) Delete(user *model.User) error {
	return m.DeleteFunc(user)
}

func (m *MockUserRepository) Login(email, password string) (string, error) {
	if email == "test@example.com" && password == "password" {
		claims := &jwt.MapClaims{
			"user_id": 1,
			"email":   "test@example.com",
			"iss":     "Intranet API - generate by IslaIT",
			"exp":     time.Now().Add(time.Hour * 72).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token.SignedString([]byte("test_secret"))
	}
	return "", errors.New("invalid email or password")
}
