package directus

import (
	"context"
	"fmt"

	"github.com/rhyoharianja/go-directusSDK/client"
)

// Alteration represents a schema alteration
type Alteration struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Collection string                 `json:"collection"`
	Field      string                 `json:"field,omitempty"`
	Meta       map[string]interface{} `json:"meta,omitempty"`
	Schema     map[string]interface{} `json:"schema,omitempty"`
}

// AlterationService provides methods for managing schema alterations
type AlterationService struct {
	client *client.HTTPClient
}

// newAlterationService creates a new alteration service
func newAlterationService(client *client.HTTPClient) *AlterationService {
	return &AlterationService{client: client}
}

// Get retrieves all alterations
func (a *AlterationService) Get(ctx context.Context) ([]Alteration, error) {
	var alterations []Alteration
	if err := a.client.Get(ctx, "/schema/alterations", nil, &alterations); err != nil {
		return nil, fmt.Errorf("get alterations failed: %w", err)
	}
	return alterations, nil
}

// Apply applies a schema alteration
func (a *AlterationService) Apply(ctx context.Context, alteration *Alteration) error {
	return a.client.Post(ctx, "/schema/alterations", alteration, nil)
}

// Revert reverts a schema alteration
func (a *AlterationService) Revert(ctx context.Context, id string) error {
	return a.client.Delete(ctx, fmt.Sprintf("/schema/alterations/%s", id))
}
