package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rhyoharianja/go-directusSDK"
	"github.com/rhyoharianja/go-directusSDK/client"
	"github.com/rhyoharianja/go-directusSDK/models"
)

func main() {
	// Create client with custom options
	client := directus.NewClient(
		"https://your-directus-instance.com",
		client.WithTimeout(60*time.Second),
		client.WithDebug(),
	)
	
	// Example 1: Complex filtering
	fmt.Println("=== Complex Filtering Example ===")
	itemsRepo := client.Items("products")
	
	// Complex filter with nested conditions
	filter := map[string]interface{}{
		"_and": []map[string]interface{}{
			{
				"status": map[string]interface{}{
					"_eq": "published",
				},
			},
			{
				"price": map[string]interface{}{
					"_gte": 100,
					"_lte": 1000,
				},
			},
			{
				"category": map[string]interface{}{
					"_in": []string{"electronics", "books"},
				},
			},
		},
	}
	
	products, err := itemsRepo.GetMany(context.Background(), &models.QueryParams{
		Filter: filter,
		Sort:   []string{"-price", "name"},
		Limit:  20,
	})
	if err != nil {
		log.Fatal("Get products failed:", err)
	}
	
	fmt.Printf("Found %d products matching criteria\n", len(products))
	
	// Example 2: Deep querying
	fmt.Println("\n=== Deep Querying Example ===")
	articlesRepo := client.Items("articles")
	
	articles, err := articlesRepo.GetMany(context.Background(), &models.QueryParams{
		Fields: []string{"id", "title", "author.first_name", "author.last_name", "tags.name"},
		Filter: map[string]interface{}{
			"status": map[string]interface{}{
				"_eq": "published",
			},
		},
		Limit: 10,
	})
	if err != nil {
		log.Fatal("Get articles failed:", err)
	}
	
	fmt.Printf("Found %d articles with author info\n", len(articles))
	
	// Example 3: Batch operations
	fmt.Println("\n=== Batch Operations Example ===")
	
	// Create multiple items
	items := []interface{}{
		map[string]interface{}{
			"title": "Product 1",
			"price": 99.99,
			"stock": 100,
		},
		map[string]interface{}{
			"title": "Product 2",
			"price": 149.99,
			"stock": 50,
		},
	}
	
	for _, item := range items {
		_, err := itemsRepo.Create(context.Background(), item)
		if err != nil {
			log.Printf("Failed to create item: %v", err)
		}
	}
	
	fmt.Println("Created multiple products")
	
	// Example 4: File upload
	fmt.Println("\n=== File Upload Example ===")
	fileRepo := client.Files()
	
	// This would typically read from a file
	fileData := &models.FileUploadRequest{
		Filename:    "example.txt",
		Title:       "Example File",
		Description: "This is an example file",
		File:        []byte("Hello, Directus!"),
	}
	
	uploadedFile, err := fileRepo.Upload(context.Background(), fileData)
	if err != nil {
		log.Fatal("File upload failed:", err)
	}
	
	fmt.Printf("Uploaded file with ID: %s\n", uploadedFile.ID)
	
	// Example 5: User management
	fmt.Println("\n=== User Management Example ===")
	
	newUser := &models.User{
		Email:     "newuser@example.com",
		FirstName: "John",
		LastName:  "Doe",
		Role:      "editor",
	}
	
	createdUser, err := client.Users().Create(context.Background(), newUser)
	if err != nil {
		log.Fatal("Create user failed:", err)
	}
	
	fmt.Printf("Created user with ID: %s\n", createdUser.ID)
	
	// Update user
	createdUser.FirstName = "Jane"
	updatedUser, err := client.Users().Update(context.Background(), createdUser.ID, createdUser)
	if err != nil {
		log.Fatal("Update user failed:", err)
	}
	
	fmt.Printf("Updated user: %s %s\n", updatedUser.FirstName, updatedUser.LastName)
	
	// Clean up
	defer client.Close()
}
