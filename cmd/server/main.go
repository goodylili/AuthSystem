package main

import (
	"Reusable-Auth-System/pkg/api/handler"
	"Reusable-Auth-System/pkg/database"
	"Reusable-Auth-System/pkg/users"
	"fmt"
	"log"
)

// Run - is responsible for the instantiation and startup of our Go application
func Run() error {
	fmt.Println("starting up the application...")

	store, err := database.NewDatabase()
	if err != nil {
		log.Println("Database Connection Failure")
		return err
	}

	// Call the social service when you're doing this

	if err := store.MigrateDB(); err != nil {
		log.Println("failed to set up store migrations")
		return err
	}

	userService := users.NewService(store)
	httpHandler := handler.NewHandler(&userService)

	// Initialize the HTTP server
	if err := httpHandler.Initialize(); err != nil {
		log.Fatalf("Server error: %v\n", err)
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Println(err)
	}

}
