package usecase

import "learn-go-mock-test/model"

type CategoryUsecase interface {
	Get(id int) (*model.Category, error)
}
