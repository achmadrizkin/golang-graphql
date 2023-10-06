package model

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Person struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(50)"`
	Age  int    `gorm:"type:int"`
	*gorm.Model
}

type PersonWithCar struct {
	PersonId   int // GORM will use "person_id" as the column name
	PersonName string
	PersonAge  int
	CarName    string
	CarColor   string
}

func DefinePersonWithCarType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "PersonWithCar",
		Fields: graphql.Fields{
			"PersonId":   &graphql.Field{Type: graphql.Int},
			"PersonName": &graphql.Field{Type: graphql.String},
			"PersonAge":  &graphql.Field{Type: graphql.Int},
			"CarName":    &graphql.Field{Type: graphql.String},
			"CarColor":   &graphql.Field{Type: graphql.String},
		},
	})
}

func DefinePersonType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Person",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.Int},
			"name": &graphql.Field{Type: graphql.String},
			"age":  &graphql.Field{Type: graphql.Int},
		},
	})
}
