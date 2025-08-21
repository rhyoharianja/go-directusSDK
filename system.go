package directus

import (
	"context"
	"fmt"
)

// SystemService handles system configuration and info operations
type SystemService struct {
	client *Client
}

// NewSystemService creates a new system service
func NewSystemService(client *Client) *SystemService {
	return &SystemService{client: client}
}

// GetInfo retrieves system information
func (s *SystemService) GetInfo(ctx context.Context) (*SystemInfo, error) {
	var resp struct {
		Data SystemInfo `json:"data"`
	}
	path := "/system/info"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get system info: %s", response.Status())
	}

	return &resp.Data, nil
}

// GetSettings retrieves system settings
func (s *SystemService) GetSettings(ctx context.Context) (*SystemSettings, error) {
	var resp struct {
		Data SystemSettings `json:"data"`
	}
	path := "/system/settings"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get system settings: %s", response.Status())
	}

	return &resp.Data, nil
}

// UpdateSettings updates system settings
func (s *SystemService) UpdateSettings(ctx context.Context, settings *SystemSettings) (*SystemSettings, error) {
	var resp struct {
		Data SystemSettings `json:"data"`
	}
	path := "/system/settings"

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(settings).
		SetResult(&resp).
		Patch(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to update system settings: %s", response.Status())
	}

	return &resp.Data, nil
}

// SystemInfo represents system information
type SystemInfo struct {
	ID          string   `json:"id,omitempty"`
	ProjectName string   `json:"project_name"`
	ProjectURL  string   `json:"project_url"`
	Version     string   `json:"version"`
	Extensions  []string `json:"extensions,omitempty"`
}

// SystemSettings represents system settings
type SystemSettings struct {
	ID                    string                   `json:"id,omitempty"`
	ProjectName           string                   `json:"project_name"`
	ProjectURL            string                   `json:"project_url"`
	DefaultLanguage       string                   `json:"default_language"`
	DefaultTimezone       string                   `json:"default_timezone"`
	AuthLoginAttempts     int                      `json:"auth_login_attempts"`
	AuthPasswordPolicy    string                   `json:"auth_password_policy"`
	StorageAssetTransform string                   `json:"storage_asset_transform"`
	StorageAssetPresets   []map[string]interface{} `json:"storage_asset_presets,omitempty"`
	CustomCSS             string                   `json:"custom_css,omitempty"`
	ModuleBar             []map[string]interface{} `json:"module_bar,omitempty"`
}
