package directus

import (
	"context"
	"fmt"
)

// FlowService handles flow automation operations
type FlowService struct {
	client *Client
}

// NewFlowService creates a new flow service
func NewFlowService(client *Client) *FlowService {
	return &FlowService{client: client}
}

// List retrieves all flows
func (s *FlowService) List(ctx context.Context) ([]Flow, error) {
	var resp struct {
		Data []Flow `json:"data"`
	}
	path := "/flows"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to list flows: %s", response.Status())
	}

	return resp.Data, nil
}

// Get retrieves a flow by ID
func (s *FlowService) Get(ctx context.Context, id string) (*Flow, error) {
	var resp struct {
		Data Flow `json:"data"`
	}
	path := fmt.Sprintf("/flows/%s", id)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get flow: %s", response.Status())
	}

	return &resp.Data, nil
}

// Create creates a new flow
func (s *FlowService) Create(ctx context.Context, flow *Flow) (*Flow, error) {
	var resp struct {
		Data Flow `json:"data"`
	}
	path := "/flows"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(flow).
		SetResult(&resp).
		Post(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to create flow: %s", response.Status())
	}

	return &resp.Data, nil
}

// Update updates an existing flow
func (s *FlowService) Update(ctx context.Context, id string, flow *Flow) (*Flow, error) {
	var resp struct {
		Data Flow `json:"data"`
	}
	path := fmt.Sprintf("/flows/%s", id)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(flow).
		SetResult(&resp).
		Patch(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to update flow: %s", response.Status())
	}

	return &resp.Data, nil
}

// Delete deletes a flow
func (s *FlowService) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("/flows/%s", id)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		Delete(path)

	if err != nil {
		return err
	}

	if response.StatusCode() != 204 {
		return fmt.Errorf("failed to delete flow: %s", response.Status())
	}

	return nil
}

// Trigger triggers a flow
func (s *FlowService) Trigger(ctx context.Context, id string, payload map[string]interface{}) error {
	path := fmt.Sprintf("/flows/%s/trigger", id)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(payload).
		Post(path)

	if err != nil {
		return err
	}

	if response.StatusCode() != 200 {
		return fmt.Errorf("failed to trigger flow: %s", response.Status())
	}

	return nil
}
