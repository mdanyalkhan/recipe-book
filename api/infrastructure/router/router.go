package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdanyalkhan/recipe-book/api/interactors/controllers"
)

func NewRouter(gorillaRouter *mux.Router, c controllers.AppController) *mux.Router {
	gorillaRouter.HandleFunc("/", c.GetRecipes).Methods(http.MethodGet)
	gorillaRouter.HandleFunc("/", c.AddRecipe).Methods(http.MethodPost)
	gorillaRouter.HandleFunc("/", c.UpdateRecipe).Methods(http.MethodPut)

	gorillaRouter.HandleFunc("/{id:[0-9]+}", c.GetRecipe).Methods(http.MethodGet)
	gorillaRouter.HandleFunc("/{id:[0-9]+}", c.DeleteRecipe).Methods(http.MethodDelete)

	return gorillaRouter
}
