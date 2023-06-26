package main

import (
	"School_gql/config"
	"School_gql/graph"
	"School_gql/graph/generated"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "database":
			config.MigrateDB(config.GetDB())
			log.Print("\n Database loaded...")
			os.Exit(1)
		default:
			log.Print("\n Starting server....")
		}

	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
