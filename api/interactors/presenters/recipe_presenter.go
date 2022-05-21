package presenters

import "github.com/mdanyalkhan/recipe-book/api/models"

type recipePresenter struct{}

type RecipePresenter interface {
	ResponseRecipe(r *models.Recipe) *models.Recipe
}

func NewRecipePresenter() *recipePresenter {
	return &recipePresenter{}
}

func (r *recipePresenter) ResponseRecipe(recipe *models.Recipe) *models.Recipe {
	return recipe
}