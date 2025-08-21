package directus

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

// Config represents the configuration for the Directus client
type Config struct {
	BaseURL  string
	Token    string
	Email    string
	Password string
	Timeout  time.Duration
}

// Client represents a Directus API client
type Client struct {
	httpClient  *resty.Client
	baseURL     string
	token       string
	Collections *CollectionsService
	Items       *ItemsService
	Files       *FilesService
	Users       *UsersService
	Roles       *RolesService
	Services    *ServicesService
	System      *SystemService
	Settings    *SettingsService
	Flow        *FlowService
	Relations   *RelationsService
}

// NewClient creates a new Directus client
func NewClient(config Config) (*Client, error) {
	if config.BaseURL == "" {
		return nil, fmt.Errorf("base URL is required")
	}

	client := &Client{
		baseURL: config.BaseURL,
		token:   config.Token,
	}

	// Initialize HTTP client
	httpClient := resty.New().
		SetBaseURL(config.BaseURL).
		SetTimeout(config.Timeout)

	// Set authentication
	if config.Token != "" {
		httpClient.SetAuthToken(config.Token)
	} else if config.Email != "" && config.Password != "" {
		// Authenticate with email/password
		token, err := authenticate(httpClient, config.Email, config.Password)
		if err != nil {
			return nil, fmt.Errorf("authentication failed: %w", err)
		}
		client.token = token
		httpClient.SetAuthToken(token)
	}

	client.httpClient = httpClient

	// Initialize services
	client.Collections = NewCollectionsService(client)
	client.Items = NewItemsService(client)
	client.Files = NewFilesService(client)
	client.Users = NewUsersService(client)
	client.Roles = NewRolesService(client)
	client.Services = NewServicesService(client)
	client.System = NewSystemService(client)
	client.Settings = NewSettingsService(client)
	client.Flow = NewFlowService(client)
	client.Relations = NewRelationsService(client)

	return client, nil
}

// GetBaseURL returns the base URL of the client
func (c *Client) GetBaseURL() string {
	return c.baseURL
}

// GetToken returns the current token
func (c *Client) GetToken() string {
	return c.token
}

// authenticate performs authentication with email/password
func authenticate(client *resty.Client, email, password string) (string, error) {
	var resp struct {
		Data struct {
			AccessToken string `json:"access_token"`
		} `json:"data"`
	}

	response, err := client.R().
		SetBody(map[string]string{
			"email":    email,
			"password": password,
		}).
		SetResult(&resp).
		Post("/auth/login")

	if err != nil {
		return "", err
	}

	if response.StatusCode() != 200 {
		return "", fmt.Errorf("authentication failed: %s", response.Status())
	}

	return resp.Data.AccessToken, nil
}
