package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mdanyalkhan/recipe-book/api/data"
)

type Recipes struct {
	log *log.Logger
	db  *sql.DB
}

func NewRecipes(log *log.Logger, db *sql.DB) *Recipes {
	return &Recipes{log: log, db: db}
}

func (recipes *Recipes) GetRecipes(rw http.ResponseWriter, r *http.Request) {
	recipes.log.Println("Handle GET recipes")
	recipeList, err := data.GetRecipes(recipes.db)
	if err != nil {
		recipes.log.Println(err)
		http.Error(rw, "Unable to fetch from db", http.StatusInternalServerError)
	}
	err = recipeList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (recipes *Recipes) GetRecipe(rw http.ResponseWriter, r *http.Request) {
	recipes.log.Println("Handle GET recipe")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
	}
	recipe, err := data.GetRecipe(recipes.db, id)
	if err != nil {
		http.Error(rw, "Unable to find recipe", http.StatusNotFound)
	}

	err = recipe.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
