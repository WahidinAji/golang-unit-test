package service

import (
	"errors"
	"unit-testing/entity"
	"unit-testing/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

//add method to CategoryRepository

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}