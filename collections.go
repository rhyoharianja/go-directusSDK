package directus

import (
	"context"
	"fmt"
)

// CollectionsService handles operations on collections
type CollectionsService struct {
	client *Client
}

// NewCollectionsService creates a new collections service
func NewCollectionsService(client *Client) *CollectionsService {
	return &CollectionsService{client: client}
}

// Get retrieves a collection by name
func (s *CollectionsService) Get(ctx context.Context, name string) (*Collection, error) {
	var resp struct {
		Data Collection `json:"data"`
	}
	path := fmt.Sprintf("/collections/%s", name)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get collection: %s", response.Status())
	}

	return &resp.Data, nil
}

// List retrieves all collections
func (s *CollectionsService) List(ctx context.Context) ([]Collection, error) {
	var resp struct {
		Data []Collection `json:"data"`
	}
	path := "/collections"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to list collections: %s", response.Status())
	}

	return resp.Data, nil
}

// Create creates a new collection
func (s *CollectionsService) Create(ctx context.Context, collection *Collection) (*Collection, error) {
	var resp struct {
		Data Collection `json:"data"`
	}
	path := "/collections"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(collection).
		SetResult(&resp).
		Post(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to create collection: %s", response.Status())
	}

	return &resp.Data, nil
}

// Update updates an existing collection
func (s *CollectionsService) Update(ctx context.Context, name string, collection *Collection) (*Collection, error) {
	var resp struct {
		Data Collection `json:"data"`
	}
	path := fmt.Sprintf("/collections/%s", name)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(collection).
		SetResult(&resp).
		Patch(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to update collection: %s", response.Status())
	}

	return &resp.Data, nil
}

// Delete deletes a collection
func (s *CollectionsService) Delete(ctx context.Context, name string) error {
	path := fmt.Sprintf("/collections/%s", name)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		Delete(path)

	if err != nil {
		return err
	}

	if response.StatusCode() != 204 {
		return fmt.Errorf("failed to delete collection: %s", response.Status())
	}

	return nil
}
