package model

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

type Car struct {
	Id        int    `gorm:"type:int;primary_key"`
	Person_Id int    `gorm:"type:int"`
	Name      string `gorm:"type:varchar(50)"`
	Color     string `gorm:"type:varchar(50)"`
	*gorm.Model
}

func DefineCarType() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "create_car",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.Int},
			"person_id": &graphql.Field{Type: graphql.Int},
			"name":      &graphql.Field{Type: graphql.String},
			"color":     &graphql.Field{Type: graphql.String},
		},
	})
}
