package db

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
	"gorm.io/gorm"
)

type menuTreeRepository struct {
	db *gorm.DB
}

func NewMenuTreeRepository(db *gorm.DB) repository.MenuTreeRepository {
	return &menuTreeRepository{db}
}

func (r *menuTreeRepository) CreateOrUpdate(menuTree *model.MenuTree) error {
	if menuTree.ID != 0 {
		return r.db.Save(menuTree).Error
	}
	return r.db.Create(menuTree).Error
}

func (r *menuTreeRepository) GetAll() ([]*model.MenuTree, error) {
	var menuTrees []*model.MenuTree
	if err := r.db.Find(&menuTrees).Error; err != nil {
		return nil, err
	}
	return menuTrees, nil
}

func (r *menuTreeRepository) Paginate(page int, pageSize int) ([]*model.MenuTree, int, error) {
	var menuTrees []*model.MenuTree
	var total int64

	r.db.Model(&model.MenuTree{}).Count(&total)

	offset := (page - 1) * pageSize
	if err := r.db.Limit(pageSize).Offset(offset).Find(&menuTrees).Error; err != nil {
		return nil, 0, err
	}

	return menuTrees, int(total), nil
}

func (r *menuTreeRepository) Delete(menuTree *model.MenuTree) error {
	return r.db.Delete(menuTree).Error
}
