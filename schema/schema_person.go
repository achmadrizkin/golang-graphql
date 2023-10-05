package schema

import (
	"go-graphql-2/domain"
	"go-graphql-2/model"

	"github.com/graphql-go/graphql"
)

type PersonSchema struct {
	personUseCase domain.PersonUseCase
}

func NewPersonSchema(personUseCase domain.PersonUseCase) *PersonSchema {
	return &PersonSchema{personUseCase: personUseCase}
}

func (a *PersonSchema) DefineQueryPersonType(personType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// You can add a dummy field here if you don't have actual query fields
			"dummy": &graphql.Field{
				Type: graphql.String, // You can use any type here
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello, world!", nil // You can return any value
				},
			},
			"getAllPerson": &graphql.Field{
				Type:    graphql.NewList(personType), // Use graphql.NewList to specify a list type
				Resolve: a.getAllPersonResolver(),
			},
		},
	})
}

func (a *PersonSchema) DefineMutationPersonType(personType *graphql.Object) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createPerson": &graphql.Field{
				Type: personType,
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"age": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: a.createPersonResolver(), // Change this line
			},
		},
	})
}

func (a *PersonSchema) getAllPersonResolver() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		personData, err := a.personUseCase.GetAllPerson()
		if err != nil {
			return personData, err
		}

		return personData, nil
	}
}

func (a *PersonSchema) createPersonResolver() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		age, _ := p.Args["age"].(int)

		// Create a new Person and save it to the database
		// add new usecase
		person := model.Person{Name: name, Age: age}
		personData, err := a.personUseCase.CreatePerson(person)
		if err != nil {
			return personData, err
		}

		return person, nil
	}
}
