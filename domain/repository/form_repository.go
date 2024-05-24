package repository

import "github.com/drossan/core-api/domain/model"

type FormRepository interface {
	CreateOrUpdate(form *model.Form) error
	GetAll() ([]*model.Form, error)
	Paginate(page int, pageSize int) ([]*model.Form, int, error)
	Delete(form *model.Form) error
}
