package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rhyoharianja/go-directusSDK/errors"
)

// HTTPClient is the main HTTP client for Directus API
type HTTPClient struct {
	client *resty.Client
}

// Option defines configuration options for the client
type Option func(*HTTPClient)

// NewHTTPClient creates a new HTTP client
func NewHTTPClient(baseURL string, opts ...Option) *HTTPClient {
	client := resty.New().
		SetBaseURL(baseURL).
		SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetTimeout(30 * time.Second)

	httpClient := &HTTPClient{
		client: client,
	}

	for _, opt := range opts {
		opt(httpClient)
	}

	return httpClient
}

// WithTimeout sets the timeout for requests
func WithTimeout(timeout time.Duration) Option {
	return func(c *HTTPClient) {
		c.client.SetTimeout(timeout)
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *HTTPClient) {
		c.client.SetHTTPClient(httpClient)
	}
}

// WithDebug enables debug logging
func WithDebug() Option {
	return func(c *HTTPClient) {
		c.client.SetDebug(true)
	}
}

// Get performs a GET request
func (c *HTTPClient) Get(ctx context.Context, path string, params url.Values, result interface{}) error {
	resp, err := c.client.R().
		SetContext(ctx).
		SetQueryParamsFromValues(params).
		SetResult(result).
		Get(path)

	if err != nil {
		return fmt.Errorf("GET request failed: %w", err)
	}

	return c.handleResponse(resp)
}

// Post performs a POST request
func (c *HTTPClient) Post(ctx context.Context, path string, body, result interface{}) error {
	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(result).
		Post(path)

	if err != nil {
		return fmt.Errorf("POST request failed: %w", err)
	}

	return c.handleResponse(resp)
}

// Patch performs a PATCH request
func (c *HTTPClient) Patch(ctx context.Context, path string, body, result interface{}) error {
	resp, err := c.client.R().
		SetContext(ctx).
		SetBody(body).
		SetResult(result).
		Patch(path)

	if err != nil {
		return fmt.Errorf("PATCH request failed: %w", err)
	}

	return c.handleResponse(resp)
}

// Delete performs a DELETE request
func (c *HTTPClient) Delete(ctx context.Context, path string) error {
	resp, err := c.client.R().
		SetContext(ctx).
		Delete(path)

	if err != nil {
		return fmt.Errorf("DELETE request failed: %w", err)
	}

	return c.handleResponse(resp)
}

// SetAuthToken sets the authentication token
func (c *HTTPClient) SetAuthToken(token string) {
	c.client.SetAuthToken(token)
}

// SetRefreshToken sets the refresh token
func (c *HTTPClient) SetRefreshToken(token string) {
	c.client.SetHeader("X-Refresh-Token", token)
}

// handleResponse handles the HTTP response and returns appropriate errors
func (c *HTTPClient) handleResponse(resp *resty.Response) error {
	if resp.IsSuccess() {
		return nil
	}

	// Parse error response
	var apiError errors.APIError
	if err := json.Unmarshal(resp.Body(), &apiError); err != nil {
		return &errors.APIError{
			StatusCode: resp.StatusCode(),
			Message:    resp.String(),
		}
	}

	apiError.StatusCode = resp.StatusCode()
	return &apiError
}

// Health checks the health of the Directus instance
func (c *HTTPClient) Health(ctx context.Context) error {
	_, err := c.client.R().
		SetContext(ctx).
		Get("/server/health")
	return err
}

// Close cleans up resources
func (c *HTTPClient) Close() error {
	// Resty doesn't need explicit cleanup
	return nil
}
