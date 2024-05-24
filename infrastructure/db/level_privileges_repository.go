package db

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
	"gorm.io/gorm"
)

type levelPrivilegesRepository struct {
	db *gorm.DB
}

func NewLevelPrivilegesRepository(db *gorm.DB) repository.LevelPrivilegesRepository {
	return &levelPrivilegesRepository{db}
}

func (r *levelPrivilegesRepository) CreateOrUpdate(levelPrivileges *model.LevelPrivileges) error {
	if levelPrivileges.ID != 0 {
		return r.db.Save(levelPrivileges).Error
	}
	return r.db.Create(levelPrivileges).Error
}

func (r *levelPrivilegesRepository) GetAll() ([]*model.LevelPrivileges, error) {
	var levelPrivileges []*model.LevelPrivileges
	if err := r.db.Find(&levelPrivileges).Error; err != nil {
		return nil, err
	}
	return levelPrivileges, nil
}

func (r *levelPrivilegesRepository) Delete(levelPrivileges *model.LevelPrivileges) error {
	return r.db.Delete(levelPrivileges).Error
}
