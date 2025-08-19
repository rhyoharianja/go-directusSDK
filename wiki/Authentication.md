# Authentication Methods

The Go Directus SDK supports multiple authentication methods.

## 🔐 Email/Password Authentication

```go
client := directus.NewClient("https://your-instance.com", nil)
auth := client.Auth()

// Login with email/password
token, err := auth.Login(context.Background(), "admin@example.com", "password")
if err != nil {
    log.Fatal(err)
}
```

## 🔑 Static Token Authentication

```go
client := directus.NewClient("https://your-instance.com", &directus.Config{
    Token: "your-static-token",
})
```

## 🌐 OAuth Authentication

```go
client := directus.NewClient("https://your-instance.com", &directus.Config{
    Provider: "github", // or google, facebook, etc.
    Token:    "oauth-access-token",
})
```

## 🔄 Token Refresh

The SDK automatically handles token refresh when using email/password authentication.

## 📋 Token Management

```go
// Get current token
currentToken := auth.GetToken()

// Check if authenticated
isAuthenticated := auth.IsAuthenticated()

// Logout
err := auth.Logout(context.Background())
