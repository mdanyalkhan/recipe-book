package data

import (
	"database/sql"
	"encoding/json"
	"io"
)

type Recipe struct {
	RecipeSummary
	Ingredients  []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

type RecipeSummary struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RecipeSummaries []*RecipeSummary

func GetRecipes(db *sql.DB) (RecipeSummaries, error) {
	rows, err := db.Query("SELECT ID, Name, Description FROM recipes")
	if err != nil {
		return nil, err
	}
	recipeSummaries := RecipeSummaries{}
	defer rows.Close()
	for rows.Next() {
		var r RecipeSummary
		if err := rows.Scan(&r.ID, &r.Name, &r.Description); err != nil {
			return nil, err
		}
		recipeSummaries = append(recipeSummaries, &r)
	}
	return recipeSummaries, nil
}

func (recipeSummaries *RecipeSummaries) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(recipeSummaries)
}

func GetRecipe(db *sql.DB, id int) (*Recipe, error) {
	var r Recipe
	var instructions []string
	var ingredients []string

	// Read recipe summary
	err := db.QueryRow("SELECT id, name, description FROM recipes WHERE recipes.id=$1", id).Scan(&r.ID, &r.Description, &r.Name)
	if err != nil {
		return nil, err
	}

	// Read recipe instructions
	instructionRows, err := db.Query("SELECT instruction FROM recipe_instructions WHERE recipe_instructions.recipe_id=$1 ORDER BY recipe_instructions.step ASC", id)
	if err != nil {
		return nil, err
	}
	defer instructionRows.Close()
	for instructionRows.Next() {
		var instruction string
		instructionRows.Scan(&instruction)
		instructions = append(instructions, instruction)
	}

	// Read recipe ingredients
	ingredientsRow, err := db.Query("SELECT instruction FROM recipe_instructions WHERE recipe_instructions.recipe_id=$1 ORDER BY recipe_instructions.step ASC", id)
	if err != nil {
		return nil, err
	}
	defer ingredientsRow.Close()
	for ingredientsRow.Next() {
		var ingredient string
		ingredientsRow.Scan(&ingredient)
		ingredients = append(ingredients, ingredient)
	}

	r.Ingredients = ingredients
	r.Instructions = instructions
	return &r, nil
}

func (recipe *Recipe) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(recipe)
}
