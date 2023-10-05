package domain

import "go-graphql-2/model"

type CarRepo interface {
	CreateCar(model.Car) (model.Car, error)
}

type CarUseCase interface {
	CreateCar(model.Car) (model.Car, error)
}
