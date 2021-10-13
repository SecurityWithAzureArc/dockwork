package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/securitywithazurearc/dockwork/gql/server"
)

func (r *mutationResolver) RootMessage(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s", name), nil
}

func (r *queryResolver) Root(ctx context.Context) (string, error) {
	return "Hello world", nil
}

func (r *subscriptionResolver) RootNotification(ctx context.Context) (<-chan string, error) {
	ticker := time.NewTicker(1 * time.Second)
	dataChan := make(chan string)

	go func() {
		for {
			select {
			case <-ticker.C:
				dataChan <- "Hello world"
			case <-ctx.Done():
				fmt.Println("Done!")
				return
			}
		}
	}()

	return dataChan, nil
}

// Mutation returns server.MutationResolver implementation.
func (r *Resolver) Mutation() server.MutationResolver { return &mutationResolver{r} }

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

// Subscription returns server.SubscriptionResolver implementation.
func (r *Resolver) Subscription() server.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
