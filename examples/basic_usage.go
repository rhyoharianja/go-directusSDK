package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rhyoharianja/go-directusSDK"
	"github.com/rhyoharianja/go-directusSDK/utils"
)

func main() {
	// Create a new Directus client
	client := directus.NewClient("https://your-directus-instance.com")
	
	// Example 1: Authentication
	fmt.Println("=== Authentication Example ===")
	authResponse, err := client.Auth().Login(context.Background(), "admin@example.com", "password")
	if err != nil {
		log.Fatal("Login failed:", err)
	}
	
	// Set the auth token for subsequent requests
	client.SetAuthToken(authResponse.AccessToken)
	
	// Example 2: Get current user
	fmt.Println("\n=== Current User Example ===")
	user, err := client.Auth().Me(context.Background())
	if err != nil {
		log.Fatal("Get current user failed:", err)
	}
	fmt.Printf("Current user: %s %s (%s)\n", user.FirstName, user.LastName, user.Email)
	
	// Example 3: Get items with query builder
	fmt.Println("\n=== Items Example ===")
	itemsRepo := client.Items("articles")
	
	// Build query with utility
	query := utils.NewQueryBuilder().
		Fields("id", "title", "status", "published_at").
		Filter(map[string]interface{}{
			"status": map[string]interface{}{
				"_eq": "published",
			},
		}).
		Sort("-published_at").
		Limit(10).
		Build()
	
	articles, err := itemsRepo.GetMany(context.Background(), query)
	if err != nil {
		log.Fatal("Get articles failed:", err)
	}
	
	fmt.Printf("Found %d articles\n", len(articles))
	for _, article := range articles {
		fmt.Printf("- %s (ID: %s)\n", article["title"], article["id"])
	}
	
	// Example 4: Create a new item
	fmt.Println("\n=== Create Item Example ===")
	newArticle := map[string]interface{}{
		"title":       "My New Article",
		"content":     "This is the content of my new article",
		"status":      "draft",
		"published_at": time.Now(),
	}
	
	createdArticle, err := itemsRepo.Create(context.Background(), newArticle)
	if err != nil {
		log.Fatal("Create article failed:", err)
	}
	fmt.Printf("Created article with ID: %s\n", createdArticle["id"])
	
	// Example 5: Get users
	fmt.Println("\n=== Users Example ===")
	users, err := client.Users().GetMany(context.Background(), &models.QueryParams{
		Limit: 5,
		Fields: []string{"id", "email", "first_name", "last_name"},
	})
	if err != nil {
		log.Fatal("Get users failed:", err)
	}
	
	fmt.Printf("Found %d users\n", len(users))
	for _, user := range users {
		fmt.Printf("- %s %s (%s)\n", user.FirstName, user.LastName, user.Email)
	}
	
	// Example 6: Get files
	fmt.Println("\n=== Files Example ===")
	files, err := client.Files().GetMany(context.Background(), &models.QueryParams{
		Limit: 5,
		Fields: []string{"id", "filename_download", "title", "filesize"},
	})
	if err != nil {
		log.Fatal("Get files failed:", err)
	}
	
	fmt.Printf("Found %d files\n", len(files))
	for _, file := range files {
		fmt.Printf("- %s (%d bytes)\n", file.FilenameDownload, file.Filesize)
	}
	
	// Clean up
	defer client.Close()
}
