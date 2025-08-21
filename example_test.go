package directus_test

import (
	"context"
	"fmt"
	"log"

	directus "github.com/rhyoharianja/go-directusSDK"
)

func ExampleClient() {
	// Create a new Directus client
	client, err := directus.NewClient(directus.Config{
		BaseURL: "http://localhost:8055",
		Token:   "your-access-token",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Example: List collections
	collections, err := client.Collections.List(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d collections\n", len(collections))

	// Example: Get items from a collection
	items, _, err := client.Items.List(context.Background(), "articles", &directus.QueryParams{
		Limit: 10,
		Sort:  []string{"-date_created"},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d articles\n", len(items))

	// Example: Create a new item
	newItem := directus.Item{
		"title":   "Hello World",
		"content": "This is a test article",
		"status":  "published",
	}

	createdItem, err := client.Items.Create(context.Background(), "articles", newItem)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created item with ID: %v\n", createdItem["id"])
}

func ExampleAuthentication() {
	// Create client with authentication
	client, err := directus.NewClient(directus.Config{
		BaseURL:  "http://localhost:8055",
		Email:    "admin@example.com",
		Password: "password",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Authenticated with token: %s\n", client.GetToken())
}
