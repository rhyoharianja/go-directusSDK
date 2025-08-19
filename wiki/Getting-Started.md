# Getting Started with Go Directus SDK

This guide will help you get up and running with the Go Directus SDK quickly.

## 📦 Installation

```bash
go get github.com/rhyoharianja/go-directusSDK
```

## 🔧 Basic Setup

### 1. Initialize Client

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/rhyoharianja/go-directusSDK"
)

func main() {
    // Create client
    client := directus.NewClient("https://your-directus-instance.com", nil)
    
    // Authenticate
    auth := client.Auth()
    _, err := auth.Login(context.Background(), "admin@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Successfully connected to Directus!")
}
```

### 2. Configuration Options

```go
// With custom configuration
config := &directus.Config{
    Timeout: 30 * time.Second,
    Retry:   3,
}

client := directus.NewClient("https://your-instance.com", config)
```

## 🎯 Your First Query

```go
// Get all items from a collection
items, err := client.Items("articles").GetMany(context.Background(), nil)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Found %d articles\n", len(items.Data))
```

## 📚 Next Steps

- [Authentication](Authentication) - Learn about different auth methods
- [Collections](Collections) - Working with collections
- [Items](Items) - CRUD operations on items
