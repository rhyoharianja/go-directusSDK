# Go Directus SDK

A comprehensive Go SDK for Directus 11+ that provides a type-safe, modular interface for interacting with Directus APIs.

## Features

- ✅ Full Directus REST API support
- ✅ Type-safe operations with Go structs
- ✅ Repository pattern for clean architecture
- ✅ Authentication management
- ✅ File operations
- ✅ Real-time subscriptions (WebSocket)
- ✅ Comprehensive error handling
- ✅ Context support for timeouts and cancellation
- ✅ Middleware support for logging, retry, etc.

## Installation

```bash
go get github.com/rhyoharianja/go-directusSDK
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    
    "github.com/rhyoharianja/go-directusSDK"
)

func main() {
    client := directus.NewClient("https://your-directus-instance.com", nil)
    
    // Authenticate
    auth := client.Auth()
    token, err := auth.Login(context.Background(), "admin@example.com", "password")
    if err != nil {
        log.Fatal(err)
    }
    
    // Get items from a collection
    itemsRepo := client.Items("articles")
    articles, err := itemsRepo.GetMany(context.Background(), nil)
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("Found %d articles", len(articles))
}
```

## Architecture

This SDK follows a modular repository pattern with the following structure:

- `client/` - Main client and configuration
- `repositories/` - Repository implementations for different Directus resources
- `models/` - Type definitions and data structures
- `auth/` - Authentication and authorization
- `errors/` - Custom error types and handling
- `utils/` - Utility functions and helpers

## Documentation

See [docs/](docs/) for detailed API documentation and examples.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on contributing to this project.

## License

MIT License - see [LICENSE](LICENSE) for details.
