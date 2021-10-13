package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/securitywithazurearc/dockwork/gql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Image struct {
	mongo     *mongo.Collection
	listeners []chan *model.ImageInfo
}

func NewImage(mongoDB *mongo.Database) *Image {
	return &Image{
		mongo:     mongoDB.Collection("images"),
		listeners: []chan *model.ImageInfo{},
	}
}

func (s *Image) List(ctx context.Context, skip, last int) (images []*model.ImageInfo, err error) {
	res, err := s.mongo.Find(ctx, bson.M{}, options.Find().SetLimit(int64(last)).SetSkip(int64(skip)))
	if err != nil {
		return
	}

	err = res.All(ctx, &images)

	return
}

func (s *Image) Set(ctx context.Context, name, node string) (image *model.ImageInfo, err error) {
	opts := options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After)

	query := bson.M{"name": name}
	update := bson.M{
		"$addToSet":    bson.M{"nodes": node},
		"$setOnInsert": bson.M{"createdAt": time.Now()},
		"$set":         bson.M{"updatedAt": time.Now()},
	}
	res := s.mongo.FindOneAndUpdate(ctx, query, update, opts)
	err = res.Err()
	if err != nil {
		return
	}

	image = &model.ImageInfo{}
	err = res.Decode(image)
	return
}

func (s *Image) Get(ctx context.Context, name string) (image *model.ImageInfo, err error) {
	res := s.mongo.FindOne(ctx, bson.M{"name": name})
	err = res.Err()
	if err != nil {
		return
	}

	image = &model.ImageInfo{}
	err = res.Decode(image)
	return
}

func (s *Image) DeletedFromNode(ctx context.Context, name, node string) (image *model.ImageInfo, err error) {
	query := bson.M{"name": name}
	update := bson.M{"$pull": bson.M{"nodes": node}}
	res := s.mongo.FindOneAndUpdate(ctx, query, update)
	err = res.Err()
	if err != nil {
		return
	}

	image = &model.ImageInfo{}
	err = res.Decode(image)
	return
}

func (s *Image) Delete(ctx context.Context, name string) (image *model.ImageInfo, err error) {
	query := bson.M{"name": name}
	update := bson.M{"$set": bson.M{"deletedAt": time.Now()}}
	res := s.mongo.FindOneAndUpdate(ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	err = res.Err()
	if err != nil {
		return
	}

	image = &model.ImageInfo{}
	res.Decode(image)
	return
}

func (s *Image) DeleteListen(ctx context.Context, node *string) (<-chan *model.ImageInfo, error) {
	// TODO: this may be better off as a singleton for the running app instance instead of per connection
	docQuery := bson.M{"fullDocument.deletedAt": bson.M{"$exists": true}}
	if node != nil {
		docQuery["fullDocument.nodes"] = *node
	}

	query := bson.A{
		bson.M{"$match": bson.M{"$and": bson.A{
			docQuery,
			bson.M{"operationType": bson.M{"$in": bson.A{"insert", "update", "replace"}}},
		}}},
		bson.M{"$project": bson.M{"fullDocument": 1}},
	}
	stream, err := s.mongo.Watch(ctx, query, options.ChangeStream().SetFullDocument(options.UpdateLookup))
	if err != nil {
		fmt.Println("er1")
		return nil, err
	}

	infoChan := make(chan *model.ImageInfo)

	go func(infoChan chan *model.ImageInfo) {
		defer stream.Close(ctx)

		type docEvent struct {
			Document *model.ImageInfo `bson:"fullDocument"`
		}

		for stream.Next(ctx) {
			doc := docEvent{}
			if err := stream.Decode(&doc); err != nil {
				fmt.Println("error decoding document in delete watch", err)
				return
			}

			infoChan <- doc.Document
		}
	}(infoChan)

	return infoChan, nil
}
