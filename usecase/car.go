package usecase

import (
	"go-graphql-2/domain"
	"go-graphql-2/model"
)

type carUseCase struct {
	carRepo domain.CarRepo
}

func NewCarUseCase(carRepo domain.CarRepo) domain.CarUseCase {
	return &carUseCase{carRepo: carRepo}
}

// CreateCar implements domain.CarUseCase.
func (c *carUseCase) CreateCar(car model.Car) (model.Car, error) {
	return c.carRepo.CreateCar(car)
}
