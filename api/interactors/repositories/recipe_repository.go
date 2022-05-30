package repositories

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mdanyalkhan/recipe-book/api/models"
)

type recipeRepository struct {
	db *sql.DB
}

type RecipeRepository interface {
	FetchRecipes(ctx context.Context) (models.RecipeSummaries, error)
	FetchRecipe(ctx context.Context, id int) (*models.Recipe, error)
	AddNewRecipe(ctx context.Context, recipePayload models.Recipe) (int, error)
	UpdateRecipe(ctx context.Context, updatedRecipe models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, recipeId int) (int, error)
}

func NewRecipeRespository(db *sql.DB) *recipeRepository {
	return &recipeRepository{db}
}

func (r *recipeRepository) DeleteRecipe(ctx context.Context, recipeId int) (int, error) {

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return -1, err
	}
	defer tx.Rollback()
	// Delete recipe instructions
	err = deleteById(tx, "recipe_instructions", "recipe_id", recipeId)
	if err != nil {
		return -1, err
	}
	// Delete recipe ingredients
	err = deleteById(tx, "recipe_ingredients", "recipe_id", recipeId)
	if err != nil {
		return -1, err
	}
	// Delete recipe summary
	err = deleteById(tx, "recipes", "id", recipeId)
	if err != nil {
		return -1, err
	}
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return recipeId, nil
}

func (r *recipeRepository) UpdateRecipe(ctx context.Context, updatedRecipe models.Recipe) (*models.Recipe, error) {

	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Update Recipe
	_, err = tx.Exec("UPDATE recipes SET name = $1, description = $2 WHERE id = $3", updatedRecipe.Name, updatedRecipe.Description, updatedRecipe.ID)
	if err != nil {
		return nil, err
	}
	// Update recipe instructions
	err = deleteById(tx, "recipe_instructions", "recipe_id", updatedRecipe.ID)
	if err != nil {
		return nil, err
	}
	err = insertRecipeInstructions(tx, updatedRecipe.Instructions, updatedRecipe.ID)
	if err != nil {
		return nil, err
	}

	// Update recipe ingredients
	err = deleteById(tx, "recipe_ingredients", "recipe_id", updatedRecipe.ID)
	if err != nil {
		return nil, err
	}
	err = insertRecipeIngredients(tx, updatedRecipe.Ingredients, updatedRecipe.ID)
	if err != nil {
		return nil, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *recipeRepository) AddNewRecipe(ctx context.Context, recipePayload models.Recipe) (int, error) {
	tx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	// add recipe summary
	var lastInsertId int64
	stmt, err := tx.Prepare("INSERT INTO recipes(name, description) VALUES($1, $2) RETURNING id")
	if err != nil {
		log.Println(err)
		return -1, err
	}
	err = stmt.QueryRow(recipePayload.Name, recipePayload.Description).Scan(&lastInsertId)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	// Add recipe instructions
	err = insertRecipeInstructions(tx, recipePayload.Instructions, int(lastInsertId))
	if err != nil {
		return -1, err
	}

	// Add recipe ingredients
	err = insertRecipeIngredients(tx, recipePayload.Ingredients, int(lastInsertId))
	if err != nil {
		return -1, err
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return -1, err
	}
	return int(lastInsertId), nil
}

func (r *recipeRepository) FetchRecipes(ctx context.Context) (models.RecipeSummaries, error) {
	rows, err := r.db.Query("SELECT ID, Name, Description FROM recipes")
	if err != nil {
		return nil, err
	}
	recipeSummaries := models.RecipeSummaries{}
	defer rows.Close()
	counter := 0
	for rows.Next() {
		counter++
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
	defer tx.Rollback()

	// Read recipe summary
	err = tx.QueryRow("SELECT id, name, description FROM recipes WHERE recipes.id=$1", id).Scan(&recipe.ID, &recipe.Description, &recipe.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Read recipe instructions
	instructionRows, err := tx.Query("SELECT instruction FROM recipe_instructions WHERE recipe_instructions.recipe_id=$1 ORDER BY recipe_instructions.step ASC", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer instructionRows.Close()
	for instructionRows.Next() {
		var instruction string
		instructionRows.Scan(&instruction)
		instructions = append(instructions, instruction)
	}

	// Read recipe ingredients
	ingredientsRow, err := tx.Query("SELECT ingredient FROM recipe_ingredients WHERE recipe_ingredients.recipe_id=$1", id)
	if err != nil {
		log.Println(err)
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

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return &recipe, nil
}
