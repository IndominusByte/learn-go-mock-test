package dao

import "learn-go-mock-test/model"

type CategoryDao interface {
	FindById(id int) *model.Category
}
