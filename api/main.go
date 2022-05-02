package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mdanyalkhan/recipe-book/api/handlers"
	"github.com/mdanyalkhan/recipe-book/api/util"
)

const (
	sqlConfigPath = "./sql_config.yaml"
	dbname        = "recipe_book_db"
)

type App struct {
	Logging *log.Logger
	Router  *mux.Router
	Server  *http.Server
	DB      *sql.DB
}

func (app *App) InitializeRoutes() {
	recipes := handlers.NewRecipes(app.Logging, app.DB)
	getRouter := app.Router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", recipes.GetRecipe)
	getRouter.HandleFunc("/", recipes.GetRecipes)
}

func (app *App) Initialize(dbConfig util.Config) error {
	app.Logging = log.New(os.Stdout, "recipes-api", log.LstdFlags)
	app.Router = &mux.Router{}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	app.DB = db
	app.Server = &http.Server{
		Addr:         ":8888",
		Handler:      app.Router,
		ErrorLog:     app.Logging,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	app.InitializeRoutes()
	return nil
}

func (app *App) Run() error {
	app.Logging.Println("Starting server on port :8888")
	err := app.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		app.Logging.Printf("Error starting server %s\n", err)
		os.Exit(1)
	} else if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {

	dbConfig, err := util.ReadConfig(sqlConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	app := App{}
	err = app.Initialize(*dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer app.DB.Close()
	app.Logging.Println("Established connection with postgresql")

	// Start server
	go app.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got Signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	app.Server.Shutdown(ctx)
}
