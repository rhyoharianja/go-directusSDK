package directus

import (
	"context"
	"fmt"

	"github.com/rhyoharianja/go-directusSDK/client"
)

// SystemInfo represents system information
type SystemInfo struct {
	Version    string   `json:"version"`
	Extensions []string `json:"extensions"`
}

// SystemService provides methods for system operations
type SystemService struct {
	client *client.HTTPClient
}

// newSystemService creates a new system service
func newSystemService(client *client.HTTPClient) *SystemService {
	return &SystemService{client: client}
}

// GetInfo retrieves system information
func (s *SystemService) GetInfo(ctx context.Context) (*SystemInfo, error) {
	var info SystemInfo
	if err := s.client.Get(ctx, "/system/info", nil, &info); err != nil {
		return nil, fmt.Errorf("get system info failed: %w", err)
	}
	return &info, nil
}

// Reboot reboots the system
func (s *SystemService) Reboot(ctx context.Context) error {
	return s.client.Post(ctx, "/system/reboot", nil, nil)
}

// Shutdown shuts down the system
func (s *SystemService) Shutdown(ctx context.Context) error {
	return s.client.Post(ctx, "/system/shutdown", nil, nil)
}
