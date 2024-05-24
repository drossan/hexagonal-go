package db

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
	"gorm.io/gorm"
)

type levelRepository struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) repository.LevelRepository {
	return &levelRepository{db}
}

func (r *levelRepository) CreateOrUpdate(level *model.Level) error {
	if level.ID != 0 {
		return r.db.Save(level).Error
	}
	return r.db.Create(level).Error
}

func (r *levelRepository) GetByID(id uint) (*model.Level, error) {
	var level model.Level
	err := r.db.Preload("LevelPrivileges.Form").First(&level, id).Error
	return &level, err
}

func (r *levelRepository) GetAll() ([]*model.Level, error) {
	var levels []*model.Level
	if err := r.db.Find(&levels).Error; err != nil {
		return nil, err
	}
	return levels, nil
}

func (r *levelRepository) Paginate(page int, pageSize int) ([]*model.Level, int, error) {
	var levels []*model.Level
	var total int64

	r.db.Model(&model.Level{}).Count(&total)

	offset := (page - 1) * pageSize
	if err := r.db.Preload("LevelPrivileges.Form").Limit(pageSize).Offset(offset).Find(&levels).Error; err != nil {
		return nil, 0, err
	}

	return levels, int(total), nil
}

func (r *levelRepository) Delete(level *model.Level) error {
	return r.db.Delete(level).Error
}
