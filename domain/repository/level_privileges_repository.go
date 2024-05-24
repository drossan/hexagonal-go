package repository

import "github.com/drossan/core-api/domain/model"

type LevelPrivilegesRepository interface {
	CreateOrUpdate(levelPrivileges *model.LevelPrivileges) error
	GetAll() ([]*model.LevelPrivileges, error)
	Delete(levelPrivileges *model.LevelPrivileges) error
}
