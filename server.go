package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kens1n/vitechmssample/external_services"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kens1n/vitechmssample/graph"
	"github.com/kens1n/vitechmssample/graph/generated"
	"github.com/kens1n/vitechmssample/postgres"

	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

func main() {
	godotenv.Load(".env")

	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
	})

	defer db.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		HdataRepo:   postgres.HdataRepo{DB: db},
		GuidService: external_services.GuidServiceLocal{},
		HashService: external_services.HashServiceLocal{},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
