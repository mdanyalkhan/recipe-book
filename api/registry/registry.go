package registry

import (
	"database/sql"

	"github.com/mdanyalkhan/recipe-book/api/interactors/controllers"
)

type registry struct {
	db *sql.DB
}

type Registry interface {
	NewAppController() *controllers.AppController
}

func NewRegistry(db *sql.DB) *registry {
	return &registry{db}
}
func (r *registry) NewAppController() controllers.AppController {
	return r.NewRecipeController()
}
