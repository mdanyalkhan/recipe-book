package presenters

import "github.com/mdanyalkhan/recipe-book/api/models"

type recipePresenter struct{}

func NewRecipePresenter() *recipePresenter {
	return &recipePresenter{}
}

func (r *recipePresenter) ResponseRecipe(recipe *models.Recipe) *models.Recipe {
	return recipe
}

func (r *recipePresenter) ResponseRecipes(recipes models.RecipeSummaries) models.RecipeSummaries {
	return recipes
}
