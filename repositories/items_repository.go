package repositories

import (
	"context"
	"fmt"
	"net/url"

	"github.com/rhyoharianja/go-directusSDK/client"
	"github.com/rhyoharianja/go-directusSDK/models"
)

type itemsRepository struct {
	client     *client.HTTPClient
	collection string
}

// NewItemsRepository creates a new items repository
func NewItemsRepository(client *client.HTTPClient) ItemsRepository {
	return &itemsRepository{
		client: client,
	}
}

func (r *itemsRepository) WithCollection(collection string) ItemsRepository {
	return &itemsRepository{
		client:     r.client,
		collection: collection,
	}
}

func (r *itemsRepository) GetMany(ctx context.Context, params *models.QueryParams) ([]map[string]interface{}, error) {
	if r.collection == "" {
		return nil, fmt.Errorf("collection name is required")
	}

	var paramsMap url.Values
	if params != nil {
		paramsMap = url.Values{}
		if len(params.Fields) > 0 {
			for _, field := range params.Fields {
				paramsMap.Add("fields", field)
			}
		}
		if params.Limit > 0 {
			paramsMap.Add("limit", fmt.Sprintf("%d", params.Limit))
		}
		if params.Offset > 0 {
			paramsMap.Add("offset", fmt.Sprintf("%d", params.Offset))
		}
		if len(params.Sort) > 0 {
			for _, sort := range params.Sort {
				paramsMap.Add("sort", sort)
			}
		}
		if params.Search != "" {
			paramsMap.Add("search", params.Search)
		}
	}

	var response struct {
		Data []map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get items failed: %w", err)
	}

	return response.Data, nil
}

func (r *itemsRepository) GetOne(ctx context.Context, id string, params *models.QueryParams) (map[string]interface{}, error) {
	if r.collection == "" {
		return nil, fmt.Errorf("collection name is required")
	}

	var paramsMap url.Values
	if params != nil {
		paramsMap = url.Values{}
		if len(params.Fields) > 0 {
			for _, field := range params.Fields {
				paramsMap.Add("fields", field)
			}
		}
	}

	var response struct {
		Data map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s/%s", r.collection, id)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get item failed: %w", err)
	}

	return response.Data, nil
}

func (r *itemsRepository) Create(ctx context.Context, data interface{}) (map[string]interface{}, error) {
	if r.collection == "" {
		return nil, fmt.Errorf("collection name is required")
	}

	var response struct {
		Data map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Post(ctx, path, data, &response); err != nil {
		return nil, fmt.Errorf("create item failed: %w", err)
	}

	return response.Data, nil
}

func (r *itemsRepository) Update(ctx context.Context, id string, data interface{}) (map[string]interface{}, error) {
	if r.collection == "" {
		return nil, fmt.Errorf("collection name is required")
	}

	var response struct {
		Data map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s/%s", r.collection, id)
	if err := r.client.Patch(ctx, path, data, &response); err != nil {
		return nil, fmt.Errorf("update item failed: %w", err)
	}

	return response.Data, nil
}

func (r *itemsRepository) Delete(ctx context.Context, id string) error {
	if r.collection == "" {
		return fmt.Errorf("collection name is required")
	}

	path := fmt.Sprintf("/items/%s/%s", r.collection, id)
	if err := r.client.Delete(ctx, path); err != nil {
		return fmt.Errorf("delete item failed: %w", err)
	}

	return nil
}

func (r *itemsRepository) DeleteMany(ctx context.Context, ids []string) error {
	if r.collection == "" {
		return fmt.Errorf("collection name is required")
	}

	data := map[string]interface{}{
		"keys": ids,
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Delete(ctx, path); err != nil {
		return fmt.Errorf("delete items failed: %w", err)
	}

	return nil
}
