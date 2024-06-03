package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/harshitw/goGraphql/graph"
)

const defaultPort = "8080"

// http://localhost:8080/  - Tool like postman built in for graphql

// schema, resolvers, query, mutations

// Under the hood query/mutation will be put intp HTTP POST request and
// it's a text format protocol as it will be the payload

/*
For list of users, give list of todos
Resolvers - Function calls that get data using database
You'll go to get all users from db, get todos from db, link them and package it together using resolvers

	query {
	  users {
	    todos {
	      id
	      name
	      status
	    }
	  }
	}
*/
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
