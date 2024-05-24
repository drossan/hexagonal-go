package db

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
	"github.com/drossan/core-api/helpers"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.Preload("Level.LevelPrivileges.Form").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Paginate(page int, pageSize int) ([]*model.User, int, error) {
	var users []*model.User
	var total int64

	r.db.Model(&model.User{}).Count(&total)

	offset := (page - 1) * pageSize
	if err := r.db.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, int(total), nil
}

func (r *userRepository) Delete(user *model.User) error {
	return r.db.Delete(user).Error
}

func (r *userRepository) Login(email, password string) (string, error) {
	user, err := r.GetByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	ps := sha256.Sum256([]byte(password))
	pwd := fmt.Sprintf("%x", ps)

	if pwd == user.Password {
		token, err := helpers.GenerateJWT(user)
		if err != nil {
			return "", err
		}

		return token, nil

	}

	return "", errors.New("invalid email or password")
}
