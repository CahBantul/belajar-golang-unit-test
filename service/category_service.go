package service

import (
	"errors"

	"github.com/CahBantul/belajar-golang-unit-test/entity"
	"github.com/CahBantul/belajar-golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error)  {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("category Not Found")
	}else{
		return category, nil
	}
}