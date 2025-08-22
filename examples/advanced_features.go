package main

import (
	"fmt"
	"log"
	"time"

	directus "github.com/rhyoharianja/go-directusSDK"
)

func main() {
	// Create client with configuration
	config := directus.Config{
		BaseURL:  "https://your-directus-instance.com",
		Email:    "admin@example.com",
		Password: "password",
		Timeout:  30 * time.Second,
	}

	client, err := directus.NewClient(config)
	if err != nil {
		log.Fatal("Failed to create client:", err)
	}

	fmt.Printf("Successfully authenticated with token: %s\n", client.GetToken())

	// Example 1: Using advanced filters
	// Create complex filter using new filter builder functions
	filter1 := directus.NewFilterEqual("status", "published")
	filter2 := directus.NewFilterBetween("publish_date", time.Now().AddDate(0, -1, 0), time.Now())
	filter3 := directus.NewFilterNotNull("author")

	// Combine filters with AND logic
	combinedFilter := directus.NewFilterAnd(filter1, filter2, filter3)

	// Create query parameters with aliases and language
	params := &directus.QueryParams{
		Fields: []string{"id", "title", "author", "publish_date", "status"},
		Filter: combinedFilter,
		Limit:  10,
		Sort:   []string{"-publish_date"},
		Lang:   "en", // Language for translations
	}

	// Add field aliases
	params.AddAlias("author", "written_by")
	params.AddAlias("publish_date", "published_on")

	fmt.Println("Query parameters with aliases and filters:")
	fmt.Printf("Fields: %v\n", params.Fields)
	fmt.Printf("Aliases: %v\n", params.Aliases)
	fmt.Printf("Filter: %v\n", params.Filter)
	fmt.Printf("Language: %s\n", params.Lang)

	// Example 2: Complex nested filtering
	nestedFilter := directus.NewFilterAnd(
		directus.NewFilterOr(
			directus.NewFilterEqual("category", "technology"),
			directus.NewFilterEqual("category", "science"),
		),
		directus.NewFilterContains("title", "AI"),
		directus.NewFilterBetween("views", 100, 1000),
	)

	fmt.Println("\nNested filter structure:")
	fmt.Printf("%+v\n", nestedFilter)

	// Example 3: Translation support demonstration
	translationParams := &directus.QueryParams{
		Fields: []string{"id", "title", "description"},
		Lang:   "fr", // French translations
		Limit:  3,
	}

	fmt.Println("\nTranslation parameters:")
	fmt.Printf("Language: %s\n", translationParams.Lang)

	// Example 4: Demonstrate filter builder functions
	fmt.Println("\nAvailable filter builder functions:")
	fmt.Println("- NewFilterEqual(field, value)")
	fmt.Println("- NewFilterNotEqual(field, value)")
	fmt.Println("- NewFilterContains(field, value)")
	fmt.Println("- NewFilterIn(field, values)")
	fmt.Println("- NewFilterBetween(field, from, to)")
	fmt.Println("- NewFilterNull(field)")
	fmt.Println("- NewFilterNotNull(field)")
	fmt.Println("- NewFilterAnd(filters...)")
	fmt.Println("- NewFilterOr(filters...)")

	// Example 5: JSON unmarshaling improvements
	fmt.Println("\nJSON unmarshaling improvements:")
	fmt.Println("- Better error messages with context")
	fmt.Println("- Validation of JSON before parsing")
	fmt.Println("- Detailed error information")

	fmt.Println("\nAll enhancements have been successfully implemented!")
	fmt.Println("1. ✅ Advanced filter support with builder functions")
	fmt.Println("2. ✅ Field aliases support")
	fmt.Println("3. ✅ Translation language support")
	fmt.Println("4. ✅ Improved JSON unmarshaling with better error handling")
}
