package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ImageInfo struct {
	ID    primitive.ObjectID `json:"id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Nodes []ImageInfoNode    `json:"nodes" bson:"nodes"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt" bson:"deletedAt"`
}

func (i ImageNodeInput) ToImageInfoNode() ImageInfoNode {
	return ImageInfoNode(i)
}
