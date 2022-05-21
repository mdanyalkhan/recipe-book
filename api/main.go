package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/mdanyalkhan/recipe-book/api/infrastructure/datastore"
	"github.com/mdanyalkhan/recipe-book/api/infrastructure/router"
	"github.com/mdanyalkhan/recipe-book/api/registry"
	"github.com/mdanyalkhan/recipe-book/api/util"
)

const configPath = "./sql_config.yaml"

func main() {
	config, err := util.ReadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	db := datastore.NewDB(*config)
	defer db.Close()
	registry := registry.NewRegistry(db)
	appController := registry.NewAppController()
	r := &mux.Router{}
	r = router.NewRouter(r, appController)

	server := &http.Server{Addr: ":8888", Handler: r}

	go func() {
		log.Println("Starting server on port :8888")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Error starting server %s\n", err)
			os.Exit(1)
		} else if err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Go signal: ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
