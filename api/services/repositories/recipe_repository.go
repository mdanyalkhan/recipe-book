package repositories

import (
	"context"

	"github.com/mdanyalkhan/recipe-book/api/models"
)

type RecipeRepository interface {
	FetchRecipes(ctx context.Context) (models.RecipeSummaries, error)
	FetchRecipe(ctx context.Context, id int) (*models.Recipe, error)
	AddNewRecipe(ctx context.Context, recipePayload models.Recipe) (int, error)
	UpdateRecipe(ctx context.Context, recipePayload models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, recipeId int) (int, error)
}
