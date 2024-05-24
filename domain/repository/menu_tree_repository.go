package repository

import "github.com/drossan/core-api/domain/model"

type MenuTreeRepository interface {
	CreateOrUpdate(menuTree *model.MenuTree) error
	GetAll() ([]*model.MenuTree, error)
	Paginate(page int, pageSize int) ([]*model.MenuTree, int, error)
	Delete(menuTree *model.MenuTree) error
}
