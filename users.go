package directus

import (
	"context"
	"fmt"
)

// UsersService handles user operations
type UsersService struct {
	client *Client
}

// NewUsersService creates a new users service
func NewUsersService(client *Client) *UsersService {
	return &UsersService{client: client}
}

// Get retrieves a user by ID
func (s *UsersService) Get(ctx context.Context, id string) (*User, error) {
	var resp struct {
		Data User `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/users/%s", id))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get user: %s", response.Status())
	}

	return &resp.Data, nil
}

// List retrieves all users
func (s *UsersService) List(ctx context.Context, params *QueryParams) ([]User, error) {
	var resp struct {
		Data []User `json:"data"`
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
		if params.Offset > 0 {
			req.SetQueryParam("offset", fmt.Sprintf("%d", params.Offset))
		}
		if len(params.Sort) > 0 {
			req.SetQueryParam("sort", joinFields(params.Sort))
		}
	}

	response, err := req.Get("/users")
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to list users: %s", response.Status())
	}

	return resp.Data, nil
}

// Create creates a new user
func (s *UsersService) Create(ctx context.Context, user *User) (*User, error) {
	var resp struct {
		Data User `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(user).
		SetResult(&resp).
		Post("/users")

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to create user: %s", response.Status())
	}

	return &resp.Data, nil
}

// Update updates an existing user
func (s *UsersService) Update(ctx context.Context, id string, user *User) (*User, error) {
	var resp struct {
		Data User `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(user).
		SetResult(&resp).
		Patch(fmt.Sprintf("/users/%s", id))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to update user: %s", response.Status())
	}

	return &resp.Data, nil
}

// Delete deletes a user
func (s *UsersService) Delete(ctx context.Context, id string) error {
	response, err := s.client.httpClient.R().
		SetContext(ctx).
		Delete(fmt.Sprintf("/users/%s", id))

	if err != nil {
		return err
	}

	if response.StatusCode() != 204 {
		return fmt.Errorf("failed to delete user: %s", response.Status())
	}

	return nil
}
