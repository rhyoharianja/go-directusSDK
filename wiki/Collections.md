# Working with Collections

## List Collections
```go
collections, err := client.Collections().List(context.Background())
```

## Get Collection Schema
```go
schema, err := client.Collections().GetSchema(context.Background(), "articles")
```

## Create Collection
```go
collection := &directus.Collection{
    Collection: "products",
    Meta: &directus.CollectionMeta{
        Hidden: false,
        Icon:   "shopping_cart",
    },
    Schema: &directus.CollectionSchema{
        Name: "Products",
    },
}

err := client.Collections().Create(context.Background(), collection)
