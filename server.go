package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	// "github.com/yangwawa0323/gqlgen-todo3/graph"
	"github.com/yangwawa0323/gqlgen-todo3/graph/generated"
	"github.com/yangwawa0323/gqlgen-todo3/models"
	"github.com/yangwawa0323/gqlgen-todo3/resolvers"
)

const defaultPort = "8080"

var (
	globalResolver *resolvers.Resolver
)

func init() {
	globalResolver = &resolvers.Resolver{
		Chapters: &[]models.Chapter{
			{Title: "Chapter 1"},
			{Title: "Chapter 2"},
		},
		Classes: &[]models.Class{},
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: globalResolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
