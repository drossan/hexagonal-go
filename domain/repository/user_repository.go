package repository

import "github.com/drossan/core-api/domain/model"

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Paginate(page int, pageSize int) ([]*model.User, int, error)
	Delete(user *model.User) error
	Login(email, password string) (string, error)
}
