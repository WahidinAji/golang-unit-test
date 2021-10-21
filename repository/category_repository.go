package repository

import "unit-testing/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

//after this, we create service of repository.