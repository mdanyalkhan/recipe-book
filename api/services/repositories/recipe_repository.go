package repositories

import (
	"context"

	"github.com/mdanyalkhan/recipe-book/api/models"
)

type RecipeRepository interface {
	FetchRecipe(ctx context.Context, id int) (*models.Recipe, error)
}
