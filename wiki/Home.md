# Go Directus SDK Wiki

Welcome to the official wiki for the Go Directus SDK! This wiki provides comprehensive documentation, tutorials, and examples for using the SDK effectively.

## 📚 Quick Navigation

- [**Getting Started**](Getting-Started) - Installation and basic usage
- [**Authentication**](Authentication) - All authentication methods
- [**Collections**](Collections) - Working with Directus collections
- [**Items**](Items) - CRUD operations on items
- [**Querying**](Querying) - Advanced querying techniques
- [**Files**](Files) - File operations and management
- [**Users & Roles**](Users-and-Roles) - User and role management
- [**Examples**](Examples) - Complete working examples
- [**API Reference**](API-Reference) - Detailed API documentation
- [**Troubleshooting**](Troubleshooting) - Common issues and solutions

## 🚀 Quick Start

```go
import "github.com/rhyoharianja/go-directusSDK"

client := directus.NewClient("https://your-instance.com", nil)
auth := client.Auth()
_, err := auth.Login(context.Background(), "admin@example.com", "password")
```

## 📖 Latest Updates

- **v1.0.0** - Initial release with full Directus 11+ support
- **Enhanced Querying** - Advanced filtering and aliasing
- **File Operations** - Complete file management support
- **Real-time Support** - WebSocket subscriptions

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](https://github.com/rhyoharianja/go-directusSDK/blob/main/CONTRIBUTING.md) for details.

## 📞 Support

- [Issues](https://github.com/rhyoharianja/go-directusSDK/issues)
- [Discussions](https://github.com/rhyoharianja/go-directusSDK/discussions)
