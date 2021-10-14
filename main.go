package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-nm/sig"
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

	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	router.Handle("/graphql", gql.Handler(mongoClient.Database(cfg.DatabaseName)))
	if cfg.GraphQLPlayEnabled {
		router.Handle("/", playground.Handler("Dock Work", "/graphql"))
	}

	server := http.Server{Addr: cfg.Addr, Handler: router}
	return sig.StopSignalE(server.ListenAndServe, func() error {
		return server.Shutdown(context.TODO())
	})
}

func main() {
	config.Load()

	if err := runAPI(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run the api: %s\n", err)
		os.Exit(1)
	}
}
