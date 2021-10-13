package gql

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/securitywithazurearc/dockwork/gql/resolver"
	"github.com/securitywithazurearc/dockwork/gql/server"
)

func wsOriginCheck(r *http.Request) bool {
	// TODO: set a more resonable default & allow configuration
	return true
}

func Handler(mongoDB *mongo.Database) http.Handler {
	schema := server.NewExecutableSchema(server.Config{
		Resolvers: resolver.NewResolver(mongoDB),
	})

	srv := handler.New(schema)

	srv.AddTransport(transport.Websocket{
		Upgrader:              websocket.Upgrader{CheckOrigin: wsOriginCheck},
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return srv
}
