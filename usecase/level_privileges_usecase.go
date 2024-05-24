package usecase

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
)

type LevelPrivilegesUseCase struct {
	levelPrivilegesRepository repository.LevelPrivilegesRepository
}

func NewLevelPrivilegesUseCase(levelPrivilegesRepo repository.LevelPrivilegesRepository) *LevelPrivilegesUseCase {
	return &LevelPrivilegesUseCase{levelPrivilegesRepository: levelPrivilegesRepo}
}

func (uc *LevelPrivilegesUseCase) CreateOrUpdateLevelPrivilege(levelPrivilege *model.LevelPrivileges) error {
	return uc.levelPrivilegesRepository.CreateOrUpdate(levelPrivilege)
}

func (uc *LevelPrivilegesUseCase) GetAllLevelPrivilege() ([]*model.LevelPrivileges, error) {
	return uc.levelPrivilegesRepository.GetAll()
}

func (uc *LevelPrivilegesUseCase) DeleteLevelPrivilege(levelPrivilege *model.LevelPrivileges) error {
	return uc.levelPrivilegesRepository.Delete(levelPrivilege)
}
