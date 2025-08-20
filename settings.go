package directus

import (
	"context"
	"fmt"

	"github.com/rhyoharianja/go-directusSDK/client"
)

// Settings represents Directus settings
type Settings struct {
	ID                    string        `json:"id,omitempty"`
	ProjectName           string        `json:"project_name,omitempty"`
	ProjectURL            string        `json:"project_url,omitempty"`
	ProjectColor          string        `json:"project_color,omitempty"`
	ProjectLogo           *string       `json:"project_logo,omitempty"`
	PublicForeground      *string       `json:"public_foreground,omitempty"`
	PublicBackground      *string       `json:"public_background,omitempty"`
	PublicNote            *string       `json:"public_note,omitempty"`
	AuthLoginAttempts     int           `json:"auth_login_attempts,omitempty"`
	AuthPasswordPolicy    string        `json:"auth_password_policy,omitempty"`
	StorageAssetTransform string        `json:"storage_asset_transform,omitempty"`
	StorageAssetPresets   []AssetPreset `json:"storage_asset_presets,omitempty"`
	CustomCSS             string        `json:"custom_css,omitempty"`
}

// AssetPreset represents an asset transformation preset
type AssetPreset struct {
	Key                string `json:"key"`
	Fit                string `json:"fit"`
	Width              int    `json:"width"`
	Height             int    `json:"height"`
	Quality            int    `json:"quality"`
	WithoutEnlargement bool   `json:"without_enlargement"`
}

// SettingsService provides methods for managing Directus settings
type SettingsService struct {
	client *client.HTTPClient
}

// newSettingsService creates a new settings service
func newSettingsService(client *client.HTTPClient) *SettingsService {
	return &SettingsService{client: client}
}

// Get retrieves the current settings
func (s *SettingsService) Get(ctx context.Context) (*Settings, error) {
	var settings Settings
	if err := s.client.Get(ctx, "/settings", nil, &settings); err != nil {
		return nil, fmt.Errorf("get settings failed: %w", err)
	}
	return &settings, nil
}

// Update updates the settings
func (s *SettingsService) Update(ctx context.Context, settings *Settings) (*Settings, error) {
	var updated Settings
	if err := s.client.Patch(ctx, "/settings", settings, &updated); err != nil {
		return nil, fmt.Errorf("update settings failed: %w", err)
	}
	return &updated, nil
}
