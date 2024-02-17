package service

import (
	"fmt"
	"testing"

	"github.com/CahBantul/belajar-golang-unit-test/entity"
	"github.com/CahBantul/belajar-golang-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}
func TestCategoryService_GetNotFound(t *testing.T)  {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	fmt.Println(category)
	fmt.Println(err)
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T)  {
	category := entity.Category{Id: "2", Name: "Nozami"}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	fmt.Println(result)
	fmt.Println(category)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Name, result.Name)
	assert.Equal(t, category.Id, result.Id)
}