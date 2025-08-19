# Examples

## 🚀 Basic CRUD Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rhyoharianja/go-directusSDK"
)

func main() {
    client := directus.NewClient("https://your-instance.com", nil)
    auth := client.Auth()
    
    // Authenticate
    _, err := auth.Login(context.Background(), "admin@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }

    // Create
    newArticle := map[string]interface{}{
        "title":   "My First Article",
        "content": "This is my first article using Go Directus SDK",
        "status":  "published",
    }
    
    created, err := client.Items("articles").Create(context.Background(), newArticle)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created article with ID: %v\n", created["id"])

    // Read
    article, err := client.Items("articles").GetOne(context.Background(), created["id"])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Article title: %s\n", article["title"])

    // Update
    update := map[string]interface{}{
        "title": "Updated Article Title",
    }
    
    updated, err := client.Items("articles").Update(context.Background(), created["id"], update)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Updated article: %s\n", updated["title"])

    // Delete
    err = client.Items("articles").Delete(context.Background(), created["id"])
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Article deleted successfully")
}
```

## 🔍 Advanced Querying Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rhyoharianja/go-directusSDK"
)

func main() {
    client := directus.NewClient("https://your-instance.com", nil)
    auth := client.Auth()
    
    _, err := auth.Login(context.Background(), "admin@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }

    // Complex query with all features
    query := &directus.Query{
        Fields: []string{
            "id",
            "title:article_title",
            "content:article_content",
            "published_at",
            "author.name:author_name",
            "author.email:author_email",
            "tags.tag.name:tag_names",
            "category.name:category_name",
        },
        Filter: map[string]interface{}{
            "_and": []map[string]interface{}{
                {
                    "status": map[string]interface{}{"_eq": "published"},
                },
                {
                    "published_at": map[string]interface{}{
                        "_gte": "2024-01-01",
                        "_lt":  "2024-12-31",
                    },
                },
                {
                    "category.name": map[string]interface{}{
                        "_in": []string{"Technology", "Programming"},
                    },
                },
                {
                    "tags.tag.name": map[string]interface{}{
                        "_in": []string{"Go", "Tutorial", "Best Practices"},
                    },
                },
            },
        },
        Sort:  []string{"-published_at", "title"},
        Limit: 10,
        Page:  1,
    }

    articles, err := client.Items("articles").GetMany(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d articles matching criteria\n", len(articles.Data))
    
    for _, article := range articles.Data {
        fmt.Printf("Title: %s\n", article["article_title"])
        fmt.Printf("Author: %s (%s)\n", article["author_name"], article["author_email"])
        fmt.Printf("Category: %s\n", article["category_name"])
        fmt.Printf("Tags: %v\n", article["tag_names"])
        fmt.Printf("Published: %s\n", article["published_at"])
        fmt.Println("---")
    }
}
```

## 📁 File Upload Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rhyoharianja/go-directusSDK"
)

func main() {
    client := directus.NewClient("https://your-instance.com", nil)
    auth := client.Auth()
    
    _, err := auth.Login(context.Background(), "admin@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }

    // Upload file
    file, err := client.Files().Upload(context.Background(), "/path/to/image.jpg", &directus.FileUploadOptions{
        Title:       "My Image",
        Description: "A beautiful image uploaded via Go SDK",
        Folder:      "images",
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("File uploaded successfully: %s\n", file["filename_download"])
    fmt.Printf("File ID: %v\n", file["id"])
}
```

## 👥 User Management Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rhyoharianja/go-directusSDK"
)

func main() {
    client := directus.NewClient("https://your-instance.com", nil)
    auth := client.Auth()
    
    _, err := auth.Login(context.Background(), "admin@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }

    // Create user
    newUser := &directus.User{
        Email:     "newuser@example.com",
        Password:  "securepassword123",
        FirstName: "John",
        LastName:  "Doe",
        Role:      2, // Regular user role
    }

    user, err := client.Users().Create(context.Background(), newUser)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Created user: %s\n", user["email"])

    // List users with filtering
    users, err := client.Users().List(context.Background(), &directus.Query{
        Fields: []string{"id", "email", "first_name", "last_name", "status"},
        Filter: map[string]interface{}{
            "status": map[string]interface{}{
                "_eq": "active",
            },
        },
        Sort: []string{"-created_at"},
        Limit: 10,
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d active users\n", len(users.Data))
    for _, user := range users.Data {
        fmt.Printf("User: %s %s (%s)\n", 
            user["first_name"], 
            user["last_name"], 
            user["email"])
    }
}
