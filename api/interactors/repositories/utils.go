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
