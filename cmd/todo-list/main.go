package main

import (
	"fmt"
	"os"

	"github.com/justinhbaker3/todo-list/internal/sqlstore"

	"github.com/justinhbaker3/todo-list/internal/server"
)

func main() {
	// Get the port to run the app on from the env variable
	port := os.Getenv("TODOLIST_PORT")
	// default to 8080
	if port == "" {
		port = "8080"
	}

	//store := memorystore.New()
	store, err := sqlstore.New()
	if err != nil {
		panic(err)
	}
	srv := server.New(store)

	router := srv.SetupRoutes()

	addr := fmt.Sprintf(":%s", port)
	err = router.Run(addr)
	if err != nil {
		os.Exit(1)
	}
}
