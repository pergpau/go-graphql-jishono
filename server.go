package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
	database "github.com/pergpau/go-graphql-jishono/db"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/pergpau/go-graphql-jishono/graph"
	"github.com/pergpau/go-graphql-jishono/graph/generated"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

const defaultPort = "8082"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB()

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}})))

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", port),
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(srv.ListenAndServe())
}
