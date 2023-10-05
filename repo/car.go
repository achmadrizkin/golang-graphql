package repo

import (
	"go-graphql-2/domain"
	"go-graphql-2/model"

	"gorm.io/gorm"
)

type carRepo struct {
	db *gorm.DB
}

func NewCarRepo(db *gorm.DB) domain.CarRepo {
	return &carRepo{db: db}
}

// CreateCar implements domain.CarRepo.
func (c *carRepo) CreateCar(car model.Car) (model.Car, error) {
	if err := c.db.Create(&car).Error; err != nil {
		return car, err
	}

	return car, nil
}
