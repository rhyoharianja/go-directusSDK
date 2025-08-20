package directus

import (
	"context"
	"fmt"

	"github.com/rhyoharianja/go-directusSDK/client"
)

// ServerInfo represents Directus server information
type ServerInfo struct {
	PublicURL         string `json:"public_url"`
	ProjectName       string `json:"project_name"`
	ProjectDescriptor string `json:"project_descriptor"`
	ProjectColor      string `json:"project_color"`
	ProjectLogo       string `json:"project_logo"`
	PublicForeground  string `json:"public_foreground"`
	PublicBackground  string `json:"public_background"`
	PublicNote        string `json:"public_note"`
	CustomCSS         string `json:"custom_css"`
}

// ServerHealth represents server health status
type ServerHealth struct {
	Status    string `json:"status"`
	ReleaseID string `json:"release_id"`
	ServiceID string `json:"service_id"`
	Version   string `json:"version"`
}

// ServerService provides methods for server information
type ServerService struct {
	client *client.HTTPClient
}

// newServerService creates a new server service
func newServerService(client *client.HTTPClient) *ServerService {
	return &ServerService{client: client}
}

// GetInfo retrieves server information
func (s *ServerService) GetInfo(ctx context.Context) (*ServerInfo, error) {
	var info ServerInfo
	if err := s.client.Get(ctx, "/server/info", nil, &info); err != nil {
		return nil, fmt.Errorf("get server info failed: %w", err)
	}
	return &info, nil
}

// GetHealth retrieves server health status
func (s *ServerService) GetHealth(ctx context.Context) (*ServerHealth, error) {
	var health ServerHealth
	if err := s.client.Get(ctx, "/server/health", nil, &health); err != nil {
		return nil, fmt.Errorf("get server health failed: %w", err)
	}
	return &health, nil
}

// Ping checks if the server is reachable
func (s *ServerService) Ping(ctx context.Context) error {
	return s.client.Get(ctx, "/server/ping", nil, nil)
}
