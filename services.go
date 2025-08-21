package directus

import (
	"context"
	"fmt"
)

// ServicesService handles operations on services
type ServicesService struct {
	client *Client
}

// NewServicesService creates a new services service
func NewServicesService(client *Client) *ServicesService {
	return &ServicesService{client: client}
}

// List retrieves all services
func (s *ServicesService) List(ctx context.Context) ([]Service, error) {
	var resp struct {
		Data []Service `json:"data"`
	}
	path := "/services"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to list services: %s", response.Status())
	}

	return resp.Data, nil
}

// Service represents a service in Directus
type Service struct {
	ID          string                 `json:"id,omitempty"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Status      string                 `json:"status"`
	Config      map[string]interface{} `json:"config,omitempty"`
}
