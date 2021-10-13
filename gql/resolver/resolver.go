package resolver

import (
	"github.com/securitywithazurearc/dockwork/svc"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	imageSvc *svc.Image
}

func NewResolver(mongoDB *mongo.Database) *Resolver {
	return &Resolver{
		imageSvc: svc.NewImage(mongoDB),
	}
}
