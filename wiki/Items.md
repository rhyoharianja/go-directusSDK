# Items Operations

## Basic CRUD

### Create Item
```go
newItem := map[string]interface{}{
    "title": "New Item",
    "status": "published",
}
created, err := client.Items("articles").Create(context.Background(), newItem)
```

### Read Items
```go
// Single item
item, err := client.Items("articles").GetOne(context.Background(), 123)

// Multiple items
items, err := client.Items("articles").GetMany(context.Background(), nil)
```

### Update Item
```go
update := map[string]interface{}{
    "title": "Updated Title",
}
updated, err := client.Items("articles").Update(context.Background(), 123, update)
```

### Delete Item
```go
err := client.Items("articles").Delete(context.Background(), 123)
