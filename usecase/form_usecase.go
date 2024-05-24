package usecase

import (
	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/domain/repository"
)

type FormUseCase struct {
	formRepository repository.FormRepository
}

func NewFormUseCase(formRepo repository.FormRepository) *FormUseCase {
	return &FormUseCase{formRepository: formRepo}
}

func (uc *FormUseCase) CreateOrUpdateForm(form *model.Form) error {
	return uc.formRepository.CreateOrUpdate(form)
}

func (uc *FormUseCase) GetAllForms() ([]*model.Form, error) {
	return uc.formRepository.GetAll()
}

func (uc *FormUseCase) PaginateForms(page int, pageSize int) ([]*model.Form, int, error) {
	return uc.formRepository.Paginate(page, pageSize)
}

func (uc *FormUseCase) DeleteForm(form *model.Form) error {
	return uc.formRepository.Delete(form)
}
