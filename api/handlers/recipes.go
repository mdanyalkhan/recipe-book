package handlers

import (
	"database/sql"
	"log"
	"net/http"

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
	recipeList, err := data.GetRecipesDb(recipes.db)
	if err != nil {
		recipes.log.Println(err)
		http.Error(rw, "Unable to fetch from db", http.StatusInternalServerError)
	}
	err = recipeList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
