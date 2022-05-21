package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mdanyalkhan/recipe-book/api/services/interactors"
)

type recipeController struct {
	recipeInteractor interactors.RecipeInteractor
}

type RecipeController interface {
	GetRecipe(rw http.ResponseWriter, r *http.Request)
}

func NewRecipeController(rc interactors.RecipeInteractor) *recipeController {
	return &recipeController{rc}
}

func (uc *recipeController) GetRecipe(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
		return
	}
	recipe, err := uc.recipeInteractor.Get(r.Context(), id)
	if err != nil {
		http.Error(rw, "Unable to find recipe", http.StatusBadRequest)
		return
	}
	err = recipe.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusBadRequest)
	}
}
