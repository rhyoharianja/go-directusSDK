# Advanced Querying with Go Directus SDK

This guide covers all advanced querying techniques available in the SDK.

## 🎯 Field Selection

### Basic Field Selection
```go
query := &directus.Query{
    Fields: []string{"id", "title", "content", "published_at"},
}

items, err := client.Items("articles").GetMany(context.Background(), query)
```

## 🏷️ Field Aliasing

### Basic Aliasing
```go
query := &directus.Query{
    Fields: []string{
        "id",
        "title:article_title",           // Rename 'title' to 'article_title'
        "content:body",                  // Rename 'content' to 'body'
        "published_at:publish_date",     // Rename 'published_at' to 'publish_date'
    },
}
```

### Nested Field Aliasing
```go
query := &directus.Query{
    Fields: []string{
        "id",
        "title",
        "author.first_name:author_name",
        "author.email:author_email",
        "category.name:category_name",
    },
}
```

## 🔍 Filtering

### Basic Filters
```go
// Equality
query := &directus.Query{
    Filter: map[string]interface{}{
        "status": map[string]interface{}{
            "_eq": "published",
        },
    },
}

// Contains
query := &directus.Query{
    Filter: map[string]interface{}{
        "title": map[string]interface{}{
            "_contains": "Go",
        },
    },
}

// Date range
query := &directus.Query{
    Filter: map[string]interface{}{
        "published_at": map[string]interface{}{
            "_gte": "2024-01-01",
            "_lt":  "2024-12-31",
        },
    },
}
```

### Complex Filters with AND/OR
```go
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

## 🔢 Sorting

### Basic Sorting
```go
query := &directus.Query{
    Sort: []string{"-published_at", "title"}, // - for descending
}
```

## 📊 Pagination

### Limit and Offset
```go
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

## 📈 Aggregation

### Count Items
```go
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

## 🔄 Query Builder

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

## 🎯 Complete Example

```go
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

items, err := client.Items("articles").GetMany(context.Background(), query)
