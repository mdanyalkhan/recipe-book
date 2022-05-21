package registry

import (
	"github.com/mdanyalkhan/recipe-book/api/interactors/controllers"
	"github.com/mdanyalkhan/recipe-book/api/interactors/presenters"
	"github.com/mdanyalkhan/recipe-book/api/interactors/repositories"
	"github.com/mdanyalkhan/recipe-book/api/services/interactors"
)

func (r *registry) NewRecipeController() controllers.RecipeController {
	return controllers.NewRecipeController(r.NewRecipeInteractor())
}

func (r *registry) NewRecipeInteractor() interactors.RecipeInteractor {
	return interactors.NewRecipeInteractor(r.NewRecipePresenter(), r.NewRecipeRepository())
}

func (r *registry) NewRecipePresenter() presenters.RecipePresenter {
	return presenters.NewRecipePresenter()
}

func (r *registry) NewRecipeRepository() repositories.RecipeRepository {
	return repositories.NewRecipeRespository(r.db)
}
