package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mdanyalkhan/recipe-book/api/models"
	"github.com/mdanyalkhan/recipe-book/api/services/interactors"
)

type recipeController struct {
	recipeInteractor interactors.RecipeInteractor
}

type RecipeController interface {
	GetRecipe(rw http.ResponseWriter, r *http.Request)
	GetRecipes(rw http.ResponseWriter, r *http.Request)
	AddRecipe(rw http.ResponseWriter, r *http.Request)
}

func NewRecipeController(rc interactors.RecipeInteractor) *recipeController {
	return &recipeController{rc}
}

func (rc *recipeController) GetRecipe(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
		return
	}
	recipe, err := rc.recipeInteractor.Get(r.Context(), id)
	if err != nil {
		http.Error(rw, "Unable to find recipe", http.StatusBadRequest)
		return
	}
	err = recipe.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusBadRequest)
	}
}

func (rc *recipeController) GetRecipes(rw http.ResponseWriter, r *http.Request) {
	recipes, err := rc.recipeInteractor.GetSummaries(r.Context())
	if err != nil {
		http.Error(rw, "Unable to fetch recipes", http.StatusBadRequest)
		return
	}

	err = recipes.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusBadRequest)
	}
}

func (rc *recipeController) AddRecipe(rw http.ResponseWriter, r *http.Request) {
	recipe := &models.Recipe{}
	err := recipe.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
		return
	}
	recipeId, err := rc.recipeInteractor.Add(r.Context(), *recipe)
	if err != nil {
		http.Error(rw, "Unable insert new recipe into data store", http.StatusBadRequest)
		return
	}
	json.NewEncoder(rw).Encode(map[string]int{"recipeId": recipeId})
}
