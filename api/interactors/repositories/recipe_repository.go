package repositories

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mdanyalkhan/recipe-book/api/models"
)

type recipeRepository struct {
	db *sql.DB
}

func NewRecipeRespository(db *sql.DB) *recipeRepository {
	return &recipeRepository{db}
}

func (r *recipeRepository) FetchRecipes(ctx context.Context) (models.RecipeSummaries, error) {
	rows, err := r.db.Query("SELECT ID, Name, Description FROM recipes")
	if err != nil {
		return nil, err
	}
	recipeSummaries := models.RecipeSummaries{}
	defer rows.Close()
	for rows.Next() {
		var r models.RecipeSummary
		if err := rows.Scan(&r.ID, &r.Name, &r.Description); err != nil {
			return nil, err
		}
		recipeSummaries = append(recipeSummaries, &r)
	}
	return recipeSummaries, nil
}

func (r *recipeRepository) FetchRecipe(ctx context.Context, id int) (*models.Recipe, error) {
	var recipe models.Recipe
	var instructions []string
	var ingredients []string
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	// Read recipe summary
	err = tx.QueryRow("SELECT id, name, description FROM recipes WHERE recipes.id=$1", id).Scan(&recipe.ID, &recipe.Description, &recipe.Name)
	if err != nil {
		return nil, err
	}

	// Read recipe instructions
	instructionRows, err := tx.Query("SELECT instruction FROM recipe_instructions WHERE recipe_instructions.recipe_id=$1 ORDER BY recipe_instructions.step ASC", id)
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
	ingredientsRow, err := tx.Query("SELECT instruction FROM recipe_instructions WHERE recipe_instructions.recipe_id=$1 ORDER BY recipe_instructions.step ASC", id)
	if err != nil {
		return nil, err
	}
	defer ingredientsRow.Close()
	for ingredientsRow.Next() {
		var ingredient string
		ingredientsRow.Scan(&ingredient)
		ingredients = append(ingredients, ingredient)
	}

	recipe.Ingredients = ingredients
	recipe.Instructions = instructions
	return &recipe, nil
}
