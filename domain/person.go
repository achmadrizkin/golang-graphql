package domain

import "go-graphql-2/model"

type PersonRepo interface {
	CreatePerson(model.Person) (model.Person, error)
}

type PersonUseCase interface {
	CreatePerson(model.Person) (model.Person, error)
}
