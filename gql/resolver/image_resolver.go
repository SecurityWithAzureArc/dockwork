package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/securitywithazurearc/dockwork/gql/model"
)

func (r *mutationResolver) AddImage(ctx context.Context, image model.ImageInput) (*model.ImageInfo, error) {
	return r.imageSvc.Set(ctx, image.Name, image.Node)
}

func (r *mutationResolver) DeleteImage(ctx context.Context, name string) (*model.ImageInfo, error) {
	return r.imageSvc.Delete(ctx, name)
}

func (r *mutationResolver) DeletedImage(ctx context.Context, name string, node string) (*model.ImageInfo, error) {
	return r.imageSvc.DeletedFromNode(ctx, name, node)
}

func (r *queryResolver) Images(ctx context.Context, last *int, skip *int) ([]*model.ImageInfo, error) {
	lastNum := 48
	if last != nil {
		lastNum = *last
	}
	if lastNum > 100 {
		lastNum = 100
	}

	skipNum := 0
	if skip != nil {
		skipNum = *skip
	}

	return r.imageSvc.List(ctx, skipNum, lastNum)
}

func (r *queryResolver) Image(ctx context.Context, name string) (*model.ImageInfo, error) {
	return r.imageSvc.Get(ctx, name)
}

func (r *subscriptionResolver) DeleteImageNotification(ctx context.Context, node *string) (<-chan *model.ImageInfo, error) {
	return r.imageSvc.DeleteListen(ctx, node)
}
