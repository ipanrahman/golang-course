package unittest

import "errors"

type CategoryService struct {
	Repository CategoryRepository
}

func (service *CategoryService) FindByID(id string) (*Category, error) {
	category := service.Repository.FindByID(id)
	if category == nil {
		return nil, errors.New("category not found")
	}
	return category, nil
}
