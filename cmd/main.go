package main

import (
	"fmt"

	"test-inmemory-database/internal/repository"
)

func main() {
	database := repository.NewInMemoryDatabase[string, string]()

	fmt.Println("Inserting key: 'key' with value: 'value'")
	database.Set("key", "value")
	fmt.Printf("Getting value for key: 'key' - %s\n", database.Get("key"))
	fmt.Printf("Starting transaction\n")
	database.StartTransaction()
	fmt.Println("Inserting key: 'key' with value: 'value2'")
	database.Set("key", "value2")
	fmt.Printf("Getting value for key within transaction: 'key' - %s\n", database.Get("key"))
	fmt.Printf("Committing transaction\n")
	database.Commit()
}
