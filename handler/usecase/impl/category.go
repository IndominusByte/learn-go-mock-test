package usecase_impl

import (
	"errors"
	"learn-go-mock-test/handler/dao"
	"learn-go-mock-test/model"
)

type CategoryUsecaseImpl struct {
	CategoryDao dao.CategoryDao
}

func (usecaseImpl *CategoryUsecaseImpl) Get(id int) (*model.Category, error) {
	if category := usecaseImpl.CategoryDao.FindById(id); category != nil {
		return category, nil
	}
	return nil, errors.New("Category not found!")
}
