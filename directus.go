// Package directus provides a comprehensive Go SDK for Directus 11+
package directus

import (
	"context"
	"net/http"
	"time"

	"github.com/rhyoharianja/go-directusSDK/client"
	"github.com/rhyoharianja/go-directusSDK/repositories"
)

// Client is the main entry point for the Directus SDK
type Client struct {
	httpClient *client.HTTPClient
	auth       repositories.AuthRepository
	items      repositories.ItemsRepository
	files      repositories.FilesRepository
	users      repositories.UsersRepository
	roles      repositories.RolesRepository
	collections repositories.CollectionsRepository
}

// NewClient creates a new Directus client
func NewClient(baseURL string, opts ...client.Option) *Client {
	httpClient := client.NewHTTPClient(baseURL, opts...)
	
	return &Client{
		httpClient:  httpClient,
		auth:        repositories.NewAuthRepository(httpClient),
		items:       repositories.NewItemsRepository(httpClient),
		files:       repositories.NewFilesRepository(httpClient),
		users:       repositories.NewUsersRepository(httpClient),
		roles:       repositories.NewRolesRepository(httpClient),
		collections: repositories.NewCollectionsRepository(httpClient),
	}
}

// Auth returns the authentication repository
func (c *Client) Auth() repositories.AuthRepository {
	return c.auth
}

// Items returns the items repository for a specific collection
func (c *Client) Items(collection string) repositories.ItemsRepository {
	return c.items.WithCollection(collection)
}

// Files returns the files repository
func (c *Client) Files() repositories.FilesRepository {
	return c.files
}

// Users returns the users repository
func (c *Client) Users() repositories.UsersRepository {
	return c.users
}

// Roles returns the roles repository
func (c *Client) Roles() repositories.RolesRepository {
	return c.roles
}

// Collections returns the collections repository
func (c *Client) Collections() repositories.CollectionsRepository {
	return c.collections
}

// SetAuthToken sets the authentication token for all requests
func (c *Client) SetAuthToken(token string) {
	c.httpClient.SetAuthToken(token)
}

// SetRefreshToken sets the refresh token for authentication
func (c *Client) SetRefreshToken(token string) {
	c.httpClient.SetRefreshToken(token)
}

// Close cleans up resources
func (c *Client) Close() error {
	return c.httpClient.Close()
}

// Health checks the health of the Directus instance
func (c *Client) Health(ctx context.Context) error {
	return c.httpClient.Health(ctx)
}
