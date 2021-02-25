package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/yangwawa0323/gqlgen-todo3/graph/generated"
	"github.com/yangwawa0323/gqlgen-todo3/models"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateClass(ctx context.Context, input models.NewClass) (*models.Class, error) {
	class := models.Class{
		Title:       input.Title,
		Description: input.Description,
		ClassID:     fmt.Sprintf("%d", rand.Uint64()),
	}
	*r.Resolver.Classes = append(*r.Resolver.Classes, class)
	return &class, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
