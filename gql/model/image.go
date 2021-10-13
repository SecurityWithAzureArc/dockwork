package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ImageInfo struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Nodes        []string           `json:"nodes" bson:"nodes"`
	ShouldDelete bool               `json:"shouldDelete" bson:"shouldDelete"`
}
