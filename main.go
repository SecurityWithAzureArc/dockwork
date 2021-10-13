package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/securitywithazurearc/dockwork/gql"
	"github.com/securitywithazurearc/dockwork/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func runAPI() (err error) {
	cfg := config.Get()

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(cfg.DatabaseURL))
	if err != nil {
		return
	}

	http.Handle("/graphql", gql.Handler(mongoClient.Database(cfg.DatabaseName)))
	http.Handle("/", playground.Handler("IRaaS", "/graphql"))

	// TODO: listen to os signal for shutdown
	return http.ListenAndServe(cfg.Addr, nil)
}

func main() {
	config.Load()

	if err := runAPI(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run the api: %s\n", err)
		os.Exit(1)
	}
}
