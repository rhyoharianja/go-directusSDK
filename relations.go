package directus

import (
	"context"
	"fmt"
)

// RelationsService handles relations management
type RelationsService struct {
	client *Client
}

// NewRelationsService creates a new relations service
func NewRelationsService(client *Client) *RelationsService {
	return &RelationsService{client: client}
}

// List retrieves all relations
func (s *RelationsService) List(ctx context.Context) ([]Relation, error) {
	var resp struct {
		Data []Relation `json:"data"`
	}
	path := "/relations"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to list relations: %s", response.Status())
	}

	return resp.Data, nil
}

// Get retrieves a relation by name
func (s *RelationsService) Get(ctx context.Context, name string) (*Relation, error) {
	var resp struct {
		Data Relation `json:"data"`
	}
	path := fmt.Sprintf("/relations/%s", name)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get relation: %s", response.Status())
	}

	return &resp.Data, nil
}

// Create creates a new relation
func (s *RelationsService) Create(ctx context.Context, relation *Relation) (*Relation, error) {
	var resp struct {
		Data Relation `json:"data"`
	}
	path := "/relations"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(relation).
		SetResult(&resp).
		Post(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to create relation: %s", response.Status())
	}

	return &resp.Data, nil
}

// Update updates an existing relation
func (s *RelationsService) Update(ctx context.Context, name string, relation *Relation) (*Relation, error) {
	var resp struct {
		Data Relation `json:"data"`
	}
	path := fmt.Sprintf("/relations/%s", name)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(relation).
		SetResult(&resp).
		Patch(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to update relation: %s", response.Status())
	}

	return &resp.Data, nil
}

// Delete deletes a relation
func (s *RelationsService) Delete(ctx context.Context, name string) error {
	path := fmt.Sprintf("/relations/%s", name)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		Delete(path)

	if err != nil {
		return err
	}

	if response.StatusCode() != 204 {
		return fmt.Errorf("failed to delete relation: %s", response.Status())
	}

	return nil
}

// Relation represents a relation in Directus
type Relation struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Status      string                 `json:"status"`
	Config      map[string]interface{} `json:"config,omitempty"`
}
