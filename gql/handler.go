package gql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/securitywithazurearc/dockwork/gql/resolver"
	"github.com/securitywithazurearc/dockwork/gql/server"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handler(mongoDB *mongo.Database) http.Handler {
	schema := server.NewExecutableSchema(server.Config{
		Resolvers: resolver.NewResolver(mongoDB),
	})

	return handler.NewDefaultServer(schema)
}
