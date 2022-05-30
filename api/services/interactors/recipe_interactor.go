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
	GetSummaries(ctx context.Context) (models.RecipeSummaries, error)
	Add(ctx context.Context, recipe models.Recipe) (int, error)
	Update(ctx context.Context, recipe models.Recipe) (*models.Recipe, error)
	Delete(ctx context.Context, recipeId int) (int, error)
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

func (r recipeInteractor) GetSummaries(ctx context.Context) (models.RecipeSummaries, error) {
	recipes, err := r.recipeRepository.FetchRecipes(ctx)

	if err != nil {
		return nil, err
	}

	return r.recipePresenter.ResponseRecipes(recipes), nil
}

func (r recipeInteractor) Add(ctx context.Context, recipePayload models.Recipe) (int, error) {
	recipeId, err := r.recipeRepository.AddNewRecipe(ctx, recipePayload)
	if err != nil {
		return -1, err
	}

	return r.recipePresenter.ResponseRecipeId(recipeId), nil
}

func (r recipeInteractor) Update(ctx context.Context, recipe models.Recipe) (*models.Recipe, error) {
	updatedRecipe, err := r.recipeRepository.UpdateRecipe(ctx, recipe)
	if err != nil {
		return nil, err
	}
	return r.recipePresenter.ResponseRecipe(updatedRecipe), nil
}

func (r recipeInteractor) Delete(ctx context.Context, recipeId int) (int, error) {
	_, err := r.recipeRepository.DeleteRecipe(ctx, recipeId)
	if err != nil {
		return -1, err
	}
	return r.recipePresenter.ResponseRecipeId(recipeId), nil
}
