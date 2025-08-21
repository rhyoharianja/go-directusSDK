# Directus SDK for Go

A comprehensive Go package for interacting with Directus CMS API, providing all features available in the Node.js SDK.

## Features

- ✅ Full authentication support (login, refresh, logout)
- ✅ Complete CRUD operations for items
- ✅ Collection management
- ✅ File operations (upload, download, manage)
- ✅ User management
- ✅ Role and permission management
- ✅ Query filtering, sorting, and pagination
- ✅ Type-safe responses
- ✅ Context support for cancellation
- ✅ Comprehensive error handling

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
    // Create a new Directus client
    client, err := directus.NewClient(directus.Config{
        BaseURL: "http://localhost:8055",
        Token:   "your-access-token",
    })
    if err != nil {
        log.Fatal(err)
    }

    // List collections
    collections, err := client.Collections.List(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d collections\n", len(collections))

    // Get items from a collection
    items, _, err := client.Items.List(context.Background(), "articles", &directus.QueryParams{
        Limit: 10,
        Sort:  []string{"-date_created"},
    })
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Found %d articles\n", len(items))
}
```

## Authentication

### Using Access Token
```go
client, err := directus.NewClient(directus.Config{
    BaseURL: "http://localhost:8055",
    Token:   "your-access-token",
})
```

### Using Email/Password
```go
client, err := directus.NewClient(directus.Config{
    BaseURL:  "http://localhost:8055",
    Email:    "admin@example.com",
    Password: "password",
})
```

## Usage Examples

### Items Operations
```go
// Get single item
item, err := client.Items.Get(ctx, "articles", "123")

// List items with filtering
items, meta, err := client.Items.List(ctx, "articles", &directus.QueryParams{
    Filter: map[string]interface{}{
        "status": "published",
    },
    Sort:  []string{"-date_created"},
    Limit: 10,
})

// Create new item
newItem := directus.Item{
    "title": "New Article",
    "content": "Article content",
    "status": "published",
}
created, err := client.Items.Create(ctx, "articles", newItem)

// Update item
updated, err := client.Items.Update(ctx, "articles", "123", directus.Item{
    "title": "Updated Title",
})

// Delete item
err = client.Items.Delete(ctx, "articles", "123")
```

### File Operations
```go
// Upload file
file, err := client.Files.Upload(ctx, "/path/to/file.jpg", map[string]interface{}{
    "title": "My Image",
})

// Get file
file, err := client.Files.Get(ctx, "file-id")

// List files
files, err := client.Files.List(ctx, &directus.QueryParams{
    Limit: 20,
})
```

### User Management
```go
// Create user
user := &directus.User{
    Email: "user@example.com",
    FirstName: "John",
    LastName: "Doe",
    Role: "editor",
}
createdUser, err := client.Users.Create(ctx, user)

// Invite user
err = client.Users.Invite(ctx, "newuser@example.com", "editor")
```

## API Reference

### Client
- `NewClient(config Config) (*Client, error)` - Create new client
- `GetBaseURL() string` - Get base URL
- `GetToken() string` - Get current token

### ItemsService
- `Get(ctx, collection, id string, params *QueryParams) (Item, error)`
- `List(ctx, collection string, params *QueryParams) ([]Item, *Meta, error)`
- `Create(ctx, collection string, item Item) (Item, error)`
- `Update(ctx, collection, id string, item Item) (Item, error)`
- `Delete(ctx, collection, id string) error`

### CollectionsService
- `Get(ctx, name string) (*Collection, error)`
- `List(ctx) ([]Collection, error)`
- `Create(ctx, collection *Collection) (*Collection, error)`
- `Update(ctx, name string, collection *Collection) (*Collection, error)`
- `Delete(ctx, name string) error`

### FilesService
- `Get(ctx, id string) (*File, error)`
- `List(ctx, params *QueryParams) ([]File, error)`
- `Upload(ctx, filePath string, metadata map[string]interface{}) (*File, error)`
- `Update(ctx, id string, metadata map[string]interface{}) (*File, error)`
- `Delete(ctx, id string) error`

### UsersService
- `Get(ctx, id string) (*User, error)`
- `List(ctx, params *QueryParams) ([]User, error)`
- `Create(ctx, user *User) (*User, error)`
- `Update(ctx, id string, user *User) (*User, error)`
- `Delete(ctx, id string) error`
- `Invite(ctx, email, role string) error`

### AuthService
- `Login(ctx, email, password string) (string, error)`
- `Refresh(ctx, refreshToken string) (string, error)`
- `Logout(ctx, refreshToken string) error`
- `Me(ctx) (*User, error)`

## Error Handling

All functions return errors that can be checked:

```go
item, err := client.Items.Get(ctx, "articles", "123")
if err != nil {
    // Handle error
    fmt.Printf("Error: %v\n", err)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see LICENSE file for details
