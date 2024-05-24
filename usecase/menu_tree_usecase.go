package usecase

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
)

type MenuTreeUseCase struct {
	menuTreeRepo repository.MenuTreeRepository
}

func NewMenuTreeUseCase(repo repository.MenuTreeRepository) *MenuTreeUseCase {
	return &MenuTreeUseCase{menuTreeRepo: repo}
}

func (uc *MenuTreeUseCase) CreateOrUpdateMenuTree(menu *model.MenuTree) error {
	return uc.menuTreeRepo.CreateOrUpdate(menu)
}

func (uc *MenuTreeUseCase) GetAllMenuTrees() ([]*model.MenuTree, error) {
	return uc.menuTreeRepo.GetAll()
}

func (uc *MenuTreeUseCase) PaginateMenuTrees(page int, pageSize int) ([]*model.MenuTree, int, error) {
	return uc.menuTreeRepo.Paginate(page, pageSize)
}

func (uc *MenuTreeUseCase) DeleteMenuTree(menu *model.MenuTree) error {
	return uc.menuTreeRepo.Delete(menu)
}
