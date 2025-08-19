package repositories

import (
	"context"
	"fmt"
	"net/url"

	"github.com/rhyoharianja/go-directusSDK/client"
	"github.com/rhyoharianja/go-directusSDK/models"
)

type usersRepository struct {
	client *client.HTTPClient
}

// NewUsersRepository creates a new users repository
func NewUsersRepository(client *client.HTTPClient) UsersRepository {
	return &usersRepository{
		client: client,
	}
}

func (r *usersRepository) GetMany(ctx context.Context, params *models.QueryParams) ([]*models.User, error) {
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
		Data []*models.User `json:"data"`
	}

	if err := r.client.Get(ctx, "/users", paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get users failed: %w", err)
	}

	return response.Data, nil
}

func (r *usersRepository) GetOne(ctx context.Context, id string, params *models.QueryParams) (*models.User, error) {
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
		Data *models.User `json:"data"`
	}

	path := fmt.Sprintf("/users/%s", id)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return response.Data, nil
}

func (r *usersRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	var response struct {
		Data *models.User `json:"data"`
	}

	if err := r.client.Post(ctx, "/users", user, &response); err != nil {
		return nil, fmt.Errorf("create user failed: %w", err)
	}

	return response.Data, nil
}

func (r *usersRepository) Update(ctx context.Context, id string, user *models.User) (*models.User, error) {
	var response struct {
		Data *models.User `json:"data"`
	}

	path := fmt.Sprintf("/users/%s", id)
	if err := r.client.Patch(ctx, path, user, &response); err != nil {
		return nil, fmt.Errorf("update user failed: %w", err)
	}

	return response.Data, nil
}

func (r *usersRepository) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("/users/%s", id)
	if err := r.client.Delete(ctx, path); err != nil {
		return fmt.Errorf("delete user failed: %w", err)
	}

	return nil
}
