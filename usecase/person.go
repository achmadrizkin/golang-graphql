package usecase

import (
	"go-graphql-2/domain"
	"go-graphql-2/model"
)

type personUseCase struct {
	personRepo domain.PersonRepo
}

func NewPersonUseCase(personRepo domain.PersonRepo) domain.PersonUseCase {
	return &personUseCase{personRepo: personRepo}
}

// CreatePerson implements domain.PersonUseCase.
func (p *personUseCase) CreatePerson(person model.Person) (model.Person, error) {
	return p.personRepo.CreatePerson(person)
}

// GetAllPerson implements domain.PersonUseCase.
func (p *personUseCase) GetAllPerson() ([]model.Person, error) {
	return p.personRepo.GetAllPerson()
}
