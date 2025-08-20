package repositories

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

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
		paramsMap = r.buildQueryParams(params)
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
		paramsMap = r.buildQueryParams(params)
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

func (r *itemsRepository) CreateMany(ctx context.Context, data []interface{}) ([]map[string]interface{}, error) {
	if r.collection == "" {
		return nil, fmt.Errorf("collection name is required")
	}

	var response struct {
		Data []map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Post(ctx, path, data, &response); err != nil {
		return nil, fmt.Errorf("create items failed: %w", err)
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

func (r *itemsRepository) UpdateMany(ctx context.Context, ids []string, data interface{}) ([]map[string]interface{}, error) {
	if r.collection == "" {
		return nil, fmt.Errorf("collection name is required")
	}

	var response struct {
		Data []map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	updateData := map[string]interface{}{
		"keys": ids,
		"data": data,
	}
	if err := r.client.Patch(ctx, path, updateData, &response); err != nil {
		return nil, fmt.Errorf("update items failed: %w", err)
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

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Delete(ctx, path); err != nil {
		return fmt.Errorf("delete items failed: %w", err)
	}

	return nil
}

func (r *itemsRepository) Count(ctx context.Context, filter interface{}) (int64, error) {
	if r.collection == "" {
		return 0, fmt.Errorf("collection name is required")
	}

	paramsMap := url.Values{}
	paramsMap.Add("aggregate[count]", "*")

	if filter != nil {
		if filterMap, ok := filter.(map[string]interface{}); ok {
			for key, value := range filterMap {
				paramsMap.Add(fmt.Sprintf("filter[%s]", key), fmt.Sprintf("%v", value))
			}
		}
	}

	var response struct {
		Data []struct {
			Count int64 `json:"count"`
		} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return 0, fmt.Errorf("count items failed: %w", err)
	}

	if len(response.Data) > 0 {
		return response.Data[0].Count, nil
	}
	return 0, nil
}

func (r *itemsRepository) Aggregate(ctx context.Context, operation string, field string, filter interface{}) (float64, error) {
	if r.collection == "" {
		return 0, fmt.Errorf("collection name is required")
	}

	validOperations := map[string]bool{
		"sum": true, "avg": true, "min": true, "max": true, "count": true,
	}
	if !validOperations[operation] {
		return 0, fmt.Errorf("invalid aggregate operation: %s", operation)
	}

	var paramsMap url.Values
	paramsMap = url.Values{}
	paramsMap.Add(fmt.Sprintf("aggregate[%s]", operation), field)

	if filter != nil {
		if filterMap, ok := filter.(map[string]interface{}); ok {
			for key, value := range filterMap {
				paramsMap.Add(fmt.Sprintf("filter[%s]", key), fmt.Sprintf("%v", value))
			}
		}
	}

	var response struct {
		Data []map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return 0, fmt.Errorf("aggregate items failed: %w", err)
	}

	if len(response.Data) > 0 {
		for _, item := range response.Data {
			if val, ok := item[operation].(float64); ok {
				return val, nil
			}
		}
	}
	return 0, nil
}

func (r *itemsRepository) Filter(ctx context.Context, filter map[string]interface{}, params *models.QueryParams) ([]map[string]interface{}, error) {
	if r.collection == "" {
		return nil, fmt.Errorf("collection name is required")
	}

	var paramsMap url.Values
	if params != nil {
		paramsMap = r.buildQueryParams(params)
	} else {
		paramsMap = url.Values{}
	}

	for key, value := range filter {
		paramsMap.Add(fmt.Sprintf("filter[%s]", key), fmt.Sprintf("%v", value))
	}

	var response struct {
		Data []map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return nil, fmt.Errorf("filter items failed: %w", err)
	}

	return response.Data, nil
}

func (r *itemsRepository) Relations(ctx context.Context, fields []string, params *models.QueryParams) ([]map[string]interface{}, error) {
	if r.collection == "" {
		return nil, fmt.Errorf("collection name is required")
	}

	var paramsMap url.Values
	if params != nil {
		paramsMap = r.buildQueryParams(params)
	} else {
		paramsMap = url.Values{}
	}

	if len(fields) > 0 {
		for _, field := range fields {
			paramsMap.Add("fields", field)
		}
	}

	var response struct {
		Data []map[string]interface{} `json:"data"`
	}

	path := fmt.Sprintf("/items/%s", r.collection)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get items with relations failed: %w", err)
	}

	return response.Data, nil
}

// buildQueryParams builds URL parameters from QueryParams
func (r *itemsRepository) buildQueryParams(params *models.QueryParams) url.Values {
	paramsMap := url.Values{}

	if len(params.Fields) > 0 {
		for _, field := range params.Fields {
			paramsMap.Add("fields", field)
		}
	}
	if params.Limit > 0 {
		paramsMap.Add("limit", strconv.Itoa(params.Limit))
	}
	if params.Offset > 0 {
		paramsMap.Add("offset", strconv.Itoa(params.Offset))
	}
	if len(params.Sort) > 0 {
		for _, sort := range params.Sort {
			paramsMap.Add("sort", sort)
		}
	}
	if params.Search != "" {
		paramsMap.Add("search", params.Search)
	}
	if len(params.Filter) > 0 {
		for key, value := range params.Filter {
			paramsMap.Add(fmt.Sprintf("filter[%s]", key), fmt.Sprintf("%v", value))
		}
	}

	return paramsMap
}
