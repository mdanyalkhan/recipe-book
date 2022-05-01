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
	dbname        = "testdb"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(dbConfig util.Config) error {
	app.Router = &mux.Router{}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	app.DB = db
	return nil
}
func main() {
	l := log.New(os.Stdout, "recipes-api", log.LstdFlags)

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
	l.Println("Established connection with postgresql")
	recipes := handlers.NewRecipes(l)

	router := mux.NewRouter()
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", recipes.GetRecipes)

	server := http.Server{
		Addr:         ":8888",
		Handler:      router,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start server
	go func() {
		l.Println("Starting server on port :8888")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			l.Printf("Error starting server %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Got Signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
