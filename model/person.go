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
