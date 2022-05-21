package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mdanyalkhan/recipe-book/api/interactors/controllers"
)

func NewRouter(gorillaRouter *mux.Router, c controllers.AppController) *mux.Router {
	gorillaRouter.HandleFunc("/{id:[0-9]+}", c.GetRecipe).Methods(http.MethodGet)
	return gorillaRouter
}
