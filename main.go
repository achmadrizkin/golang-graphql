package main

import (
	"go-graphql-2/config"
	"go-graphql-2/db"
	"go-graphql-2/model"
	"go-graphql-2/repo"
	"go-graphql-2/schema"
	"go-graphql-2/usecase"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	//Database
	db := db.ConnectionDB(&loadConfig)
	err = db.AutoMigrate(&model.Car{}, &model.Person{})
	if err != nil {
		log.Fatal("🚀 Could not DB Migrate", err)
	}

	log.Println("SUCCESS CONNECT DB AND MIGRATE")

	personRepo := repo.NewPersonRepo(db)
	personUseCase := usecase.NewPersonUseCase(personRepo)
	personShcema := schema.NewPersonSchema(personUseCase)

	log.Println("SUCCESS CONNECT PersonSchema")

	// GraphQL setup
	personModelGraphql := model.DefinePersonType()
	mutationType := personShcema.DefineMutationType(personModelGraphql)
	queryType := personShcema.DefineQueryType(personModelGraphql)

	log.Println("SUCCESS CONNECT GraphQL setup")

	schema, errSchema := graphql.NewSchema(graphql.SchemaConfig{
		Mutation: mutationType,
		Query:    queryType,
	})

	if errSchema != nil {
		log.Fatal("error schema: " + errSchema.Error())
	}

	log.Println("SUCCESS CONNECT GraphQL Schema")

	// Create a GraphQL handler
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true, // Enable GraphiQL for development (optional)
	})

	log.Println("SUCCESS CONNECT GraphQL Handler")

	http.Handle("/graphql", h)
	log.Println("Server is running on port 8085")
	http.ListenAndServe(":8085", nil)
}
