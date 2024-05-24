package db

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
	"gorm.io/gorm"
)

type formRepository struct {
	db *gorm.DB
}

func NewFormRepository(db *gorm.DB) repository.FormRepository {
	return &formRepository{db}
}

func (r *formRepository) CreateOrUpdate(form *model.Form) error {
	if form.ID != 0 {
		return r.db.Save(form).Error
	}
	return r.db.Create(form).Error
}

func (r *formRepository) GetByID(id uint) (*model.Form, error) {
	var form model.Form
	if err := r.db.First(&form, id).Error; err != nil {
		return nil, err
	}
	return &form, nil
}

func (r *formRepository) GetAll() ([]*model.Form, error) {
	var forms []*model.Form
	if err := r.db.Find(&forms).Error; err != nil {
		return nil, err
	}
	return forms, nil
}

func (r *formRepository) Paginate(page int, pageSize int) ([]*model.Form, int, error) {
	var forms []*model.Form
	var total int64

	r.db.Model(&model.Form{}).Count(&total)

	offset := (page - 1) * pageSize
	if err := r.db.Limit(pageSize).Offset(offset).Find(&forms).Error; err != nil {
		return nil, 0, err
	}

	return forms, int(total), nil
}

func (r *formRepository) Delete(form *model.Form) error {
	return r.db.Delete(form).Error
}
