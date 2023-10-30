package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"graphql-project/graph"
	"graphql-project/graph/model"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Printf("Intializing the database connection")
	Con, err := model.Open()
	if err != nil {
		log.Printf("database connection failed")
		return
	}
	err = model.AutoMigrate(&Con)
	if err != nil {
		log.Printf("failed to intialize tables")
		return
	}
	log.Printf("database tables created")

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
