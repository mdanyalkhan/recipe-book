package repositories

import (
	"database/sql"
	"strconv"
)

func bulkInsertValues(tx *sql.Tx, rows [][]interface{}, tableName string, columnNames []string) error {

	columns := ``
	for i, columneName := range columnNames {
		if i == 0 {
			columns = columneName
		} else {
			columns = columns + `, ` + columneName
		}
	}
	query := `INSERT INTO ` + tableName + `(` + columns + `) VALUES `
	values := []interface{}{}
	for i, row := range rows {
		values = append(values, row...)
		n := i * len(columnNames)
		query += `(`
		for j := 0; j < len(columnNames); j++ {
			query += `$` + strconv.Itoa(j+n+1) + `,`
		}
		query = query[:len(query)-1] + `),`
	}
	query = query[:len(query)-1] // remove trailing comma
	_, err := tx.Exec(query, values...)
	return err
}

func deleteById(tx *sql.Tx, tableName string, idName string, id int) error {
	_, err := tx.Exec(`DELETE FROM `+tableName+` WHERE `+idName+` = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func insertRecipeInstructions(tx *sql.Tx, recipeInstructions []string, recipeId int) error {
	rows := [][]interface{}{}
	for i, instruction := range recipeInstructions {
		rows = append(rows, []interface{}{recipeId, i + 1, instruction})
	}
	err := bulkInsertValues(tx, rows, "recipe_instructions", []string{"recipe_id", "step", "instruction"})
	if err != nil {
		return err
	}
	return nil
}

func insertRecipeIngredients(tx *sql.Tx, recipeIngredients []string, recipeId int) error {
	rows := [][]interface{}{}
	for _, ingredient := range recipeIngredients {
		rows = append(rows, []interface{}{recipeId, ingredient})
	}
	err := bulkInsertValues(tx, rows, "recipe_ingredients", []string{"recipe_id", "ingredient"})
	if err != nil {
		return err
	}
	return nil
}
