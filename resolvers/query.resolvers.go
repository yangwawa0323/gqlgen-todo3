package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/yangwawa0323/gqlgen-todo3/graph/generated"
	"github.com/yangwawa0323/gqlgen-todo3/models"
)

func (r *queryResolver) Todos(ctx context.Context) ([]models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Chapters(ctx context.Context) ([]models.Chapter, error) {
	return *r.Resolver.Chapters, nil
}

func (r *queryResolver) Classes(ctx context.Context) ([]models.Class, error) {
	return *r.Resolver.Classes, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
