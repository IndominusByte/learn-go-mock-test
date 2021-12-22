package dao_impl

import (
	"learn-go-mock-test/model"

	"github.com/stretchr/testify/mock"
)

type CategoryDaoImplMock struct {
	Mock mock.Mock
}

func (daoImpl *CategoryDaoImplMock) FindById(id int) *model.Category {
	args := daoImpl.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	category := args.Get(0).(model.Category)
	return &category
}
