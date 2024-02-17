package repository

import "github.com/CahBantul/belajar-golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}