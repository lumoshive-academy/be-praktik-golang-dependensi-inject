package main

import (
	"fmt"
	"log"
)

func main() {
	service, err := InitializMyService("Lumoshive")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(service.Greeter.Message)

	config := InitializeServiceConfig()
	fmt.Println(config.ConfigA.Message)

	// Inisialisasi dengan InMemoryStorage
	inMemoryStorage, err := InitializeCachingData()
	if err != nil {
		log.Fatal(err)
	}
	inMemoryStorage.Save("Hello, In-Memory!")
	data, _ := inMemoryStorage.Load()
	fmt.Println(data) // Output: "Hello, In-Memory!"

	// Inisialisasi dengan DatabaseStorage
	databaseStorage, err := InitializeDatabaseStorage()
	if err != nil {
		log.Fatal(err)
	}
	databaseStorage.Save("Hello, Database!")
	data, _ = databaseStorage.Load()
	fmt.Println(data) // Output: "Hello, Database!"

	notifier, err := InitializeNotifier()
	if err != nil {
		log.Fatal(err)
	}

	notifier.SendNotification("Hello, this is a notification!")
}
