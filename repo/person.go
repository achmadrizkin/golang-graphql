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

func (p *personRepo) GetAllPersonWithCar() ([]model.PersonWithCar, error) {
	var peopleWithCars []model.PersonWithCar

	if err := p.db.Table("people").
		Select("people.id AS PersonId, people.name AS person_name, people.age AS person_age, cars.name AS car_name, cars.color AS car_color").
		Joins("LEFT JOIN cars ON people.id = cars.person_id").
		Scan(&peopleWithCars).Error; err != nil {
		return nil, err
	}

	return peopleWithCars, nil
}
