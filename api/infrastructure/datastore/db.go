package datastore

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mdanyalkhan/recipe-book/api/util"
)

const (
	sqlConfigPath = "./sql_config.yaml"
	dbname        = "recipe_book_db"
)

func NewDB(config util.Config) *sql.DB {
	psqlOptions := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, dbname)
	db, err := sql.Open("postgres", psqlOptions)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
