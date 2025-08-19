package repositories

import (
	"context"
	"fmt"
	"net/url"

	"github.com/rhyoharianja/go-directusSDK/client"
	"github.com/rhyoharianja/go-directusSDK/models"
)

type rolesRepository struct {
	client *client.HTTPClient
}

// NewRolesRepository creates a new roles repository
func NewRolesRepository(client *client.HTTPClient) RolesRepository {
	return &rolesRepository{
		client: client,
	}
}

func (r *rolesRepository) GetMany(ctx context.Context, params *models.QueryParams) ([]*models.Role, error) {
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
		Data []*models.Role `json:"data"`
	}

	if err := r.client.Get(ctx, "/roles", paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get roles failed: %w", err)
	}

	return response.Data, nil
}

func (r *rolesRepository) GetOne(ctx context.Context, id string, params *models.QueryParams) (*models.Role, error) {
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
		Data *models.Role `json:"data"`
	}

	path := fmt.Sprintf("/roles/%s", id)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get role failed: %w", err)
	}

	return response.Data, nil
}

func (r *rolesRepository) Create(ctx context.Context, role *models.Role) (*models.Role, error) {
	var response struct {
		Data *models.Role `json:"data"`
	}

	if err := r.client.Post(ctx, "/roles", role, &response); err != nil {
		return nil, fmt.Errorf("create role failed: %w", err)
	}

	return response.Data, nil
}

func (r *rolesRepository) Update(ctx context.Context, id string, role *models.Role) (*models.Role, error) {
	var response struct {
		Data *models.Role `json:"data"`
	}

	path := fmt.Sprintf("/roles/%s", id)
	if err := r.client.Patch(ctx, path, role, &response); err != nil {
		return nil, fmt.Errorf("update role failed: %w", err)
	}

	return response.Data, nil
}

func (r *rolesRepository) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("/roles/%s", id)
	if err := r.client.Delete(ctx, path); err != nil {
		return fmt.Errorf("delete role failed: %w", err)
	}

	return nil
}
