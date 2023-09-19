package main

import (
	"authenticationsystem/pkg/api/handler"
	"authenticationsystem/pkg/database"
	"authenticationsystem/pkg/users"
	"fmt"
	"log"
)

// Run - is going to be responsible for / the instantiation and startup of our / go application
func Run() error {
	fmt.Println("starting up the application...")

	store, err := database.NewDatabase()
	if err != nil {
		log.Println("Database Connection Failure")
		return err
	}

	if err := store.MigrateDB(); err != nil {
		log.Println("failed to setup store migrations")
		return err
	}

	userService := users.NewStoreImpl(store)
	httpHandler := handler.NewHandler(userService)
	if err := httpHandler.Serve(); err != nil {
		log.Println("failed to gracefully serve our application")
		return err
	}

	return nil

}

func main() {
	fmt.Println("GO REST API Course")
	if err := Run(); err != nil {
		log.Println(err)
	}

}
