package presenters

import "github.com/mdanyalkhan/recipe-book/api/models"

type recipePresenter struct{}

type RecipePresenter interface {
	ResponseRecipe(r *models.Recipe) *models.Recipe
	ResponseRecipes(recipes models.RecipeSummaries) models.RecipeSummaries
	ResponseRecipeId(recipeId int) int
}

func NewRecipePresenter() *recipePresenter {
	return &recipePresenter{}
}

func (r *recipePresenter) ResponseRecipe(recipe *models.Recipe) *models.Recipe {
	return recipe
}

func (r *recipePresenter) ResponseRecipes(recipes models.RecipeSummaries) models.RecipeSummaries {
	return recipes
}

func (r *recipePresenter) ResponseRecipeId(recipeId int) int {
	return recipeId
}
