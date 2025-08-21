package main

import (
	"context"
	"fmt"
	"log"
	"time"

	directus "github.com/rhyoharianja/go-directusSDK"
)

func main() {
	// Initialize client
	client, err := directus.NewClient(directus.Config{
		BaseURL:  "http://localhost:8055",
		Email:    "admin@example.com",
		Password: "password",
		Timeout:  30 * time.Second,
	})
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}

	ctx := context.Background()

	// Example 1: List all collections
	fmt.Println("=== Collections ===")
	collections, err := client.Collections.List(ctx)
	if err != nil {
		log.Fatal("Failed to list collections:", err)
	}
	for _, col := range collections {
		fmt.Printf("- %s\n", col.Collection)
	}

	// Example 2: Create a new item
	fmt.Println("\n=== Creating Item ===")
	newArticle := directus.Item{
		"title":        "Getting Started with Directus Go SDK",
		"content":      "This is a comprehensive guide...",
		"status":       "published",
		"date_created": time.Now().Format("2006-01-02"),
	}

	createdArticle, err := client.Items.Create(ctx, "articles", newArticle)
	if err != nil {
		log.Fatal("Failed to create article:", err)
	}
	fmt.Printf("Created article with ID: %v\n", createdArticle["id"])

	// Example 3: List items with filtering
	fmt.Println("\n=== Listing Items ===")
	articles, meta, err := client.Items.List(ctx, "articles", &directus.QueryParams{
		Filter: map[string]interface{}{
			"status": map[string]string{
				"_eq": "published",
			},
		},
		Sort:   []string{"-date_created"},
		Limit:  5,
		Fields: []string{"id", "title", "status", "date_created"},
	})
	if err != nil {
		log.Fatal("Failed to list articles:", err)
	}
	fmt.Printf("Found %d published articles (total: %d)\n", len(articles), meta.TotalCount)

	// Example 4: User management
	fmt.Println("\n=== User Management ===")
	users, err := client.Users.List(ctx, &directus.QueryParams{
		Limit:  3,
		Fields: []string{"id", "email", "first_name", "last_name"},
	})
	if err != nil {
		log.Fatal("Failed to list users:", err)
	}
	for _, user := range users {
		fmt.Printf("- %s (%s %s)\n", user.Email, user.FirstName, user.LastName)
	}

	fmt.Println("\n=== SDK Demo Complete ===")
}
