package usecase_impl

import (
	"learn-go-mock-test/handler/dao"
	dao_impl "learn-go-mock-test/handler/dao/impl"
	"learn-go-mock-test/handler/usecase"
	"learn-go-mock-test/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var databaseOperation dao.CategoryDao = &dao_impl.CategoryDaoImplMock{Mock: mock.Mock{}}
var logic usecase.CategoryUsecase = &CategoryUsecaseImpl{CategoryDao: databaseOperation}

func TestCategoryNotFound(t *testing.T) {
	databaseOperation.(*dao_impl.CategoryDaoImplMock).Mock.On("FindById", 1).Return(nil)
	category, err := logic.Get(1)
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategorySuccess(t *testing.T) {
	data := model.Category{
		Id:   2,
		Name: "Nyoman Pradipta",
	}

	databaseOperation.(*dao_impl.CategoryDaoImplMock).Mock.On("FindById", 2).Return(data)

	category, err := logic.Get(2)
	assert.NotNil(t, category)
	assert.Nil(t, err)

	assert.Equal(t, data.Id, category.Id)
	assert.Equal(t, data.Name, category.Name)
}
