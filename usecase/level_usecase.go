package usecase

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
)

type LevelUseCase struct {
	levelRepository repository.LevelRepository
}

func NewLevelUseCase(levelRepo repository.LevelRepository) *LevelUseCase {
	return &LevelUseCase{levelRepository: levelRepo}
}

func (uc *LevelUseCase) CreateOrUpdateLevel(level *model.Level) error {
	return uc.levelRepository.CreateOrUpdate(level)
}

func (uc *LevelUseCase) GetLevelByID(id uint) (*model.Level, error) {
	return uc.levelRepository.GetByID(id)
}

func (uc *LevelUseCase) GetAllLevels() ([]*model.Level, error) {
	return uc.levelRepository.GetAll()
}

func (uc *LevelUseCase) PaginateLevels(page int, pageSize int) ([]*model.Level, int, error) {
	return uc.levelRepository.Paginate(page, pageSize)
}

func (uc *LevelUseCase) DeleteLevel(level *model.Level) error {
	return uc.levelRepository.Delete(level)
}
