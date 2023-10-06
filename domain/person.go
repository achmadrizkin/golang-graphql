package domain

import "go-graphql-2/model"

type PersonRepo interface {
	GetAllPerson() ([]model.Person, error)
	GetAllPersonWithCar() ([]model.PersonWithCar, error)
	CreatePerson(model.Person) (model.Person, error)
}

type PersonUseCase interface {
	GetAllPerson() ([]model.Person, error)
	GetAllPersonWithCar() ([]model.PersonWithCar, error)
	CreatePerson(model.Person) (model.Person, error)
}
