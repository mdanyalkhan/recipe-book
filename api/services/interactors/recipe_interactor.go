package interactors

import (
	"context"

	"github.com/mdanyalkhan/recipe-book/api/models"
	"github.com/mdanyalkhan/recipe-book/api/services/presenters"
	"github.com/mdanyalkhan/recipe-book/api/services/repositories"
)

type recipeInteractor struct {
	recipePresenter  presenters.RecipePresenter
	recipeRepository repositories.RecipeRepository
}

type RecipeInteractor interface {
	Get(ctx context.Context, id int) (*models.Recipe, error)
}

func NewRecipeInteractor(p presenters.RecipePresenter, r repositories.RecipeRepository) *recipeInteractor {
	return &recipeInteractor{p, r}
}

func (r recipeInteractor) Get(ctx context.Context, id int) (*models.Recipe, error) {
	recipe, err := r.recipeRepository.FetchRecipe(ctx, id)

	if err != nil {
		return nil, err
	}
	return r.recipePresenter.ResponseRecipe(recipe), nil
}
