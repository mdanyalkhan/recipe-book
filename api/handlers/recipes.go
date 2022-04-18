package handlers

import (
	"log"
	"net/http"

	"github.com/mdanyalkhan/recipe-book/api/data"
)

type Recipes struct {
	log *log.Logger
}

func NewRecipes(log *log.Logger) *Recipes {
	return &Recipes{log: log}
}

func (recipes *Recipes) GetRecipes(rw http.ResponseWriter, r *http.Request) {
	recipes.log.Println("Handle GET recipes")
	recipeList := data.GetRecipes()
	err := recipeList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
