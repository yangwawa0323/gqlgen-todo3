package resolvers

import "github.com/yangwawa0323/gqlgen-todo3/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Chapters *[]models.Chapter
	Classes  *[]models.Class
}
