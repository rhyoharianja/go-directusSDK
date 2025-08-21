package directus

import (
	"context"
	"fmt"
)

// SettingsService handles settings operations
type SettingsService struct {
	client *Client
}

// NewSettingsService creates a new settings service
func NewSettingsService(client *Client) *SettingsService {
	return &SettingsService{client: client}
}

// Settings represents the system settings in Directus
type Settings struct {
	ProjectName           string                 `json:"project_name,omitempty"`
	ProjectURL            string                 `json:"project_url,omitempty"`
	ProjectColor          string                 `json:"project_color,omitempty"`
	ProjectLogo           string                 `json:"project_logo,omitempty"`
	PublicForeground      string                 `json:"public_foreground,omitempty"`
	PublicBackground      string                 `json:"public_background,omitempty"`
	PublicNote            string                 `json:"public_note,omitempty"`
	AuthLoginAttempts     int                    `json:"auth_login_attempts,omitempty"`
	AuthPasswordPolicy    string                 `json:"auth_password_policy,omitempty"`
	StorageAssetTransform string                 `json:"storage_asset_transform,omitempty"`
	StorageAssetPresets   map[string]interface{} `json:"storage_asset_presets,omitempty"`
	CustomCSS             string                 `json:"custom_css,omitempty"`
	StorageDefault        string                 `json:"storage_default,omitempty"`
	StorageConfigured     bool                   `json:"storage_configured,omitempty"`
	Basemaps              map[string]interface{} `json:"basemaps,omitempty"`
	MapboxKey             string                 `json:"mapbox_key,omitempty"`
	ModuleBar             []interface{}          `json:"module_bar,omitempty"`
	ProjectDescriptor     string                 `json:"project_descriptor,omitempty"`
	DefaultLanguage       string                 `json:"default_language,omitempty"`
	CustomAspectRatios    []interface{}          `json:"custom_aspect_ratios,omitempty"`
}

// Get retrieves the system settings
func (s *SettingsService) Get(ctx context.Context) (*Settings, error) {
	var resp struct {
		Data Settings `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get("/settings")

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get settings: %s", response.Status())
	}

	return &resp.Data, nil
}

// Update updates the system settings
func (s *SettingsService) Update(ctx context.Context, settings *Settings) (*Settings, error) {
	var resp struct {
		Data Settings `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(settings).
		SetResult(&resp).
		Patch("/settings")

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to update settings: %s", response.Status())
	}

	return &resp.Data, nil
}
