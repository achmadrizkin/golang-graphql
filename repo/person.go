package repo

import (
	"go-graphql-2/domain"
	"go-graphql-2/model"

	"gorm.io/gorm"
)

type personRepo struct {
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) domain.PersonRepo {
	return &personRepo{db: db}
}

// GetAllPerson implements domain.PersonRepo.
func (p *personRepo) GetAllPerson() ([]model.Person, error) {
	var persons []model.Person
	if err := p.db.Find(&persons).Error; err != nil {
		return nil, err
	}

	return persons, nil
}

// CreatePerson implements domain.PersonRepo.
func (p *personRepo) CreatePerson(person model.Person) (model.Person, error) {
	if err := p.db.Create(&person).Error; err != nil {
		return person, err
	}

	return person, nil
}
