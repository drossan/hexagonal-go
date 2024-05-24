package repository

import "github.com/drossan/core-api/domain/model"

type LevelRepository interface {
	CreateOrUpdate(level *model.Level) error
	GetByID(id uint) (*model.Level, error)
	GetAll() ([]*model.Level, error)
	Paginate(page int, pageSize int) ([]*model.Level, int, error)
	Delete(level *model.Level) error
}
