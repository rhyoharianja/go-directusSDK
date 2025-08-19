package repositories

import (
	"context"
	"fmt"
	"net/url"

	"github.com/rhyoharianja/go-directusSDK/client"
	"github.com/rhyoharianja/go-directusSDK/models"
)

type collectionsRepository struct {
	client *client.HTTPClient
}

// NewCollectionsRepository creates a new collections repository
func NewCollectionsRepository(client *client.HTTPClient) CollectionsRepository {
	return &collectionsRepository{
		client: client,
	}
}

func (r *collectionsRepository) GetMany(ctx context.Context, params *models.QueryParams) ([]*models.Collection, error) {
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
		Data []*models.Collection `json:"data"`
	}

	if err := r.client.Get(ctx, "/collections", paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get collections failed: %w", err)
	}

	return response.Data, nil
}

func (r *collectionsRepository) GetOne(ctx context.Context, name string, params *models.QueryParams) (*models.Collection, error) {
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
		Data *models.Collection `json:"data"`
	}

	path := fmt.Sprintf("/collections/%s", name)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get collection failed: %w", err)
	}

	return response.Data, nil
}

func (r *collectionsRepository) Create(ctx context.Context, collection *models.Collection) (*models.Collection, error) {
	var response struct {
		Data *models.Collection `json:"data"`
	}

	if err := r.client.Post(ctx, "/collections", collection, &response); err != nil {
		return nil, fmt.Errorf("create collection failed: %w", err)
	}

	return response.Data, nil
}

func (r *collectionsRepository) Update(ctx context.Context, name string, collection *models.Collection) (*models.Collection, error) {
	var response struct {
		Data *models.Collection `json:"data"`
	}

	path := fmt.Sprintf("/collections/%s", name)
	if err := r.client.Patch(ctx, path, collection, &response); err != nil {
		return nil, fmt.Errorf("update collection failed: %w", err)
	}

	return response.Data, nil
}

func (r *collectionsRepository) Delete(ctx context.Context, name string) error {
	path := fmt.Sprintf("/collections/%s", name)
	if err := r.client.Delete(ctx, path); err != nil {
		return fmt.Errorf("delete collection failed: %w", err)
	}

	return nil
}
