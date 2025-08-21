package directus

import (
	"context"
	"fmt"
)

// RolesService handles role operations
type RolesService struct {
	client *Client
}

// NewRolesService creates a new roles service
func NewRolesService(client *Client) *RolesService {
	return &RolesService{client: client}
}

// Get retrieves a role by ID
func (s *RolesService) Get(ctx context.Context, id string) (*Role, error) {
	var resp struct {
		Data Role `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/roles/%s", id))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get role: %s", response.Status())
	}

	return &resp.Data, nil
}

// List retrieves all roles
func (s *RolesService) List(ctx context.Context, params *QueryParams) ([]Role, error) {
	var resp struct {
		Data []Role `json:"data"`
	}

	req := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp)

	if params != nil {
		if len(params.Fields) > 0 {
			req.SetQueryParam("fields", joinFields(params.Fields))
		}
		if params.Filter != nil {
			req.SetQueryParam("filter", toJSONString(params.Filter))
		}
		if params.Limit > 0 {
			req.SetQueryParam("limit", fmt.Sprintf("%d", params.Limit))
		}
	}

	response, err := req.Get("/roles")
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to list roles: %s", response.Status())
	}

	return resp.Data, nil
}

// Create creates a new role
func (s *RolesService) Create(ctx context.Context, role *Role) (*Role, error) {
	var resp struct {
		Data Role `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(role).
		SetResult(&resp).
		Post("/roles")

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to create role: %s", response.Status())
	}

	return &resp.Data, nil
}

// Update updates an existing role
func (s *RolesService) Update(ctx context.Context, id string, role *Role) (*Role, error) {
	var resp struct {
		Data Role `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(role).
		SetResult(&resp).
		Patch(fmt.Sprintf("/roles/%s", id))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to update role: %s", response.Status())
	}

	return &resp.Data, nil
}

// Delete deletes a role
func (s *RolesService) Delete(ctx context.Context, id string) error {
	response, err := s.client.httpClient.R().
		SetContext(ctx).
		Delete(fmt.Sprintf("/roles/%s", id))

	if err != nil {
		return err
	}

	if response.StatusCode() != 204 {
		return fmt.Errorf("failed to delete role: %s", response.Status())
	}

	return nil
}
