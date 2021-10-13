package model

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MarshalObjectID(oid primitive.ObjectID) graphql.Marshaler {
	return graphql.MarshalString(oid.Hex())
}

func UnmarshalObjectID(v interface{}) (primitive.ObjectID, error) {
	switch v := v.(type) {
	case primitive.ObjectID:
		return v, nil

	case string:
		return primitive.ObjectIDFromHex(v)

	default:
		return primitive.NilObjectID, fmt.Errorf("invalid format for id")
	}
}
