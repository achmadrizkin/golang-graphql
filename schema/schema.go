package schema

import (
	"go-graphql-2/domain"
	"go-graphql-2/model"

	"github.com/graphql-go/graphql"
)

type PersonSchema struct {
	personUseCase domain.PersonUseCase
	carUseCase    domain.CarUseCase
}

func NewPersonSchema(personUseCase domain.PersonUseCase, carUseCase domain.CarUseCase) *PersonSchema {
	return &PersonSchema{personUseCase: personUseCase, carUseCase: carUseCase}
}

func (a *PersonSchema) DefineQueryType(personType *graphql.Object, personWithCar *graphql.Object) *graphql.Object {
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
			"getAllPersonWithCar": &graphql.Field{
				Type:    graphql.NewList(personWithCar), // Use graphql.NewList to specify a list type
				Resolve: a.getAllPersonWithCarResolver(),
			},
		},
	})
}

func (a *PersonSchema) DefineMutationType(personType *graphql.Object, carType *graphql.Object) *graphql.Object {
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
			"createCar": &graphql.Field{
				Type: carType,
				Args: graphql.FieldConfigArgument{
					"person_id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"color": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: a.createCarResolver(), // Change this line
			},
		},
	})
}

func (a *PersonSchema) getAllPersonWithCarResolver() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		personWithCarData, err := a.personUseCase.GetAllPersonWithCar()
		if err != nil {
			return personWithCarData, err
		}

		return personWithCarData, nil
	}
}

func (a *PersonSchema) createCarResolver() graphql.FieldResolveFn {
	return func(p graphql.ResolveParams) (interface{}, error) {
		name, _ := p.Args["name"].(string)
		person_id, _ := p.Args["person_id"].(int)
		color, _ := p.Args["color"].(string)

		car := model.Car{Name: name, Color: color, Person_Id: person_id}
		carData, err := a.carUseCase.CreateCar(car)
		if err != nil {
			return carData, err
		}

		return carData, nil
	}
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
