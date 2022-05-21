package models

import (
	"encoding/json"
	"io"
)

type Recipe struct {
	RecipeSummary
	Ingredients  []string `json:"ingredients" validate:"required,len>0"`
	Instructions []string `json:"instructions" validate:"required,len>0"`
}

type RecipeSummary struct {
	ID          int    `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (recipe *Recipe) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(recipe)
}
