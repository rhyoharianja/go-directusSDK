# Go Directus SDK

A comprehensive Go client library for interacting with the [Directus](https://directus.io/) headless CMS API.

## Features

- 🚀 Full CRUD operations support
- 🔍 Advanced querying with filters, sorting, and pagination
- 🏷️ Field aliasing and selection
- 📊 Aggregation and grouping
- 🔐 Authentication management
- 📁 File upload and management
- 👥 User and role management
- 🎯 Type-safe operations with Go structs
- ⚡ High performance with HTTP/2 support

## Installation

```bash
go get github.com/rhyoharianja/go-directusSDK
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rhyoharianja/go-directusSDK"
)

func main() {
    // Initialize client
    client := directus.NewClient("https://your-directus-instance.com", nil)
    
    // Authenticate
    auth := client.Auth()
    _, err := auth.Login(context.Background(), "admin@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }

    // Get items from a collection
    itemsRepo := client.Items("articles")
    articles, err := itemsRepo.GetMany(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d articles\n", len(articles))
}
```

## Authentication

### Email/Password Authentication
```go
client := directus.NewClient("https://your-instance.com", nil)
auth := client.Auth()
token, err := auth.Login(context.Background(), "admin@example.com", "password")
```

### Static Token Authentication
```go
client := directus.NewClient("https://your-instance.com", &directus.Config{
    Token: "your-static-token",
})
```

## Collections

### List Collections
```go
collections, err := client.Collections().List(context.Background())
```

### Get Collection Schema
```go
schema, err := client.Collections().GetSchema(context.Background(), "articles")
```

## Items Operations

### Basic CRUD Operations

#### Create Item
```go
newArticle := map[string]interface{}{
    "title":   "My New Article",
    "content": "This is the content...",
    "status":  "published",
}

created, err := client.Items("articles").Create(context.Background(), newArticle)
```

#### Read Items
```go
// Get single item by ID
article, err := client.Items("articles").GetOne(context.Background(), 123)

// Get all items
articles, err := client.Items("articles").GetMany(context.Background(), nil)
```

#### Update Item
```go
update := map[string]interface{}{
    "title": "Updated Title",
    "status": "draft",
}

updated, err := client.Items("articles").Update(context.Background(), 123, update)
```

#### Delete Item
```go
err := client.Items("articles").Delete(context.Background(), 123)
```

## Advanced Querying

### Field Selection
```go
query := &directus.Query{
    Fields: []string{"id", "title", "status"},
}

items, err := client.Items("articles").GetMany(context.Background(), query)
```

### Field Aliasing
```go
query := &directus.Query{
    Fields: []string{
        "id",
        "title:article_title",        // Rename 'title' to 'article_title'
        "content:body",               // Rename 'content' to 'body'
        "author.first_name:author_name",
    },
}

items, err := client.Items("articles").GetMany(context.Background(), query)
```

### Filtering
```go
// Basic filters
query := &directus.Query{
    Filter: map[string]interface{}{
        "status": map[string]interface{}{
            "_eq": "published",
        },
        "title": map[string]interface{}{
            "_contains": "Go",
        },
    },
}

// Complex filters with AND/OR
query := &directus.Query{
    Filter: map[string]interface{}{
        "_and": []map[string]interface{}{
            {
                "status": map[string]interface{}{"_eq": "published"},
            },
            {
                "_or": []map[string]interface{}{
                    {
                        "title": map[string]interface{}{"_contains": "Go"},
                    },
                    {
                        "title": map[string]interface{}{"_contains": "Golang"},
                    },
                },
            },
        },
    },
}

// Date filters
query := &directus.Query{
    Filter: map[string]interface{}{
        "published_at": map[string]interface{}{
            "_gte": "2024-01-01",
            "_lt": "2024-12-31",
        },
    },
}

// Null checks
query := &directus.Query{
    Filter: map[string]interface{}{
        "description": map[string]interface{}{
            "_null": true,
        },
    },
}
```

### Sorting
```go
query := &directus.Query{
    Sort: []string{"-published_at", "title"}, // - for descending
}
```

### Pagination
```go
// Limit and offset
query := &directus.Query{
    Limit:  10,
    Offset: 20,
}

// Page-based pagination
query := &directus.Query{
    Page:  2,
    Limit: 10,
}
```

### Deep Filtering (Relational)
```go
// Filter by related collection
query := &directus.Query{
    Filter: map[string]interface{}{
        "author.name": map[string]interface{}{
            "_eq": "John Doe",
        },
    },
}

// Filter by many-to-many relation
query := &directus.Query{
    Filter: map[string]interface{}{
        "tags.tag.name": map[string]interface{}{
            "_in": []string{"Go", "Programming"},
        },
    },
}
```

### Aggregation
```go
// Count items
query := &directus.Query{
    Aggregate: map[string]interface{}{
        "count": "id",
    },
}

// Sum, avg, min, max
query := &directus.Query{
    Aggregate: map[string]interface{}{
        "sum":   "price",
        "avg":   "rating",
        "min":   "created_at",
        "max":   "updated_at",
    },
}
```

## Query Builder

### Using Query Builder
```go
import "github.com/rhyoharianja/go-directusSDK/utils"

// Build complex queries
qb := utils.NewQueryBuilder().
    Select("id", "title", "content").
    Where("status", "=", "published").
    Where("created_at", ">=", "2024-01-01").
    OrderBy("published_at", "desc").
    Limit(10).
    Offset(0)

query := qb.Build()
items, err := client.Items("articles").GetMany(context.Background(), query)
```

### Query Aliases
```go
// Using query aliases for complex queries
query := &directus.Query{
    Fields: []string{
        "id",
        "title:article_title",
        "content:article_content",
        "published_at",
        "author.name:author_name",
        "author.email:author_email",
        "tags.tag.name:tag_names",
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
                "tags.tag.name": map[string]interface{}{
                    "_in": []string{"Go", "Programming", "Tutorial"},
                },
            },
        },
    },
    Sort:  []string{"-published_at", "title"},
    Limit: 10,
    Page:  1,
}
```

## File Operations

### Upload File
```go
file, err := client.Files().Upload(context.Background(), "/path/to/file.jpg", &directus.FileUploadOptions{
    Title:       "My Image",
    Description: "A beautiful image",
    Folder:      "images",
})
```

### Download File
```go
fileData, err := client.Files().Download(context.Background(), fileID)
```

### List Files
```go
files, err := client.Files().List(context.Background(), &directus.Query{
    Filter: map[string]interface{}{
        "type": map[string]interface{}{
            "_eq": "image/jpeg",
        },
    },
})
```

## User Management

### Create User
```go
newUser := &directus.User{
    Email:    "newuser@example.com",
    Password: "securepassword",
    Role:     2, // Role ID
}

user, err := client.Users().Create(context.Background(), newUser)
```

### Update User
```go
update := map[string]interface{}{
    "first_name": "John",
    "last_name":  "Doe",
}

user, err := client.Users().Update(context.Background(), userID, update)
```

### List Users
```go
users, err := client.Users().List(context.Background(), &directus.Query{
    Fields: []string{"id", "email", "first_name", "last_name"},
    Filter: map[string]interface{}{
        "status": map[string]interface{}{
            "_eq": "active",
        },
    },
})
```

## Error Handling

```go
items, err := client.Items("articles").GetMany(context.Background(), query)
if err != nil {
    var apiErr *directus.APIError
    if errors.As(err, &apiErr) {
        fmt.Printf("API Error: %s (Code: %d)\n", apiErr.Message, apiErr.StatusCode)
        // Handle specific error cases
    } else {
        fmt.Printf("Unexpected error: %v\n", err)
    }
}
```

## Context Support

```go
// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

items, err := client.Items("articles").GetMany(ctx, query)

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
go func() {
    time.Sleep(2 * time.Second)
    cancel()
}()

items, err := client.Items("articles").GetMany(ctx, query)
```

## Examples

### Complete Example with All Features
```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/rhyoharianja/go-directusSDK"
)

func main() {
    // Initialize client
    client := directus.NewClient("https://your-instance.com", nil)
    
    // Authenticate
    auth := client.Auth()
    _, err := auth.Login(context.Background(), "admin@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }

    // Complex query example
    query := &directus.Query{
        Fields: []string{
            "id",
            "title",
            "content:body",
            "published_at",
            "author.name:author_name",
            "author.email:author_email",
            "tags.tag.name:tag_names",
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
                    "tags.tag.name": map[string]interface{}{
                        "_in": []string{"Go", "Programming", "Tutorial"},
                    },
                },
            },
        },
        Sort:  []string{"-published_at", "title"},
        Limit: 10,
        Page:  1,
    }

    // Execute query
    articles, err := client.Items("articles").GetMany(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Found %d articles matching criteria\n", len(articles.Data))
    
    // Process results
    for _, article := range articles.Data {
        fmt.Printf("Title: %s, Author: %s, Tags: %v\n", 
            article["title"], 
            article["author_name"], 
            article["tag_names"])
    }
}
```

## API Reference

### Client Methods

- `NewClient(baseURL string, config *Config) *Client` - Create new client
- `Auth() *AuthService` - Authentication service
- `Items(collection string) *ItemsRepository` - Items operations
- `Collections() *CollectionsService` - Collections management
- `Files() *FilesService` - File operations
- `Users() *UsersService` - User management
- `Roles() *RolesService` - Role management

### Query Methods

- `GetOne(ctx context.Context, id interface{}) (*Item, error)`
- `GetMany(ctx context.Context, query *Query) (*ItemsResponse, error)`
- `Create(ctx context.Context, data interface{}) (*Item, error)`
- `Update(ctx context.Context, id interface{}, data interface{}) (*Item, error)`
- `Delete(ctx context.Context, id interface{}) error`

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

- 📖 [Documentation](https://github.com/rhyoharianja/go-directusSDK/wiki)
- 🐛 [Issue Tracker](https://github.com/rhyoharianja/go-directusSDK/issues)
- 💬 [Discussions](https://github.com/rhyoharianja/go-directusSDK/discussions)
