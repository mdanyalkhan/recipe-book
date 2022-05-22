package presenters

import "github.com/mdanyalkhan/recipe-book/api/models"

type RecipePresenter interface {
	ResponseRecipe(recipe *models.Recipe) *models.Recipe
	ResponseRecipes(recipes models.RecipeSummaries) models.RecipeSummaries
}
