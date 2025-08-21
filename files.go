package directus

import (
	"context"
	"fmt"
	"os"
)

// FilesService handles file operations
type FilesService struct {
	client *Client
}

// NewFilesService creates a new files service
func NewFilesService(client *Client) *FilesService {
	return &FilesService{client: client}
}

// Get retrieves a file by ID
func (s *FilesService) Get(ctx context.Context, id string) (*File, error) {
	var resp struct {
		Data File `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp).
		Get(fmt.Sprintf("/files/%s", id))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get file: %s", response.Status())
	}

	return &resp.Data, nil
}

// List retrieves all files
func (s *FilesService) List(ctx context.Context, params *QueryParams) ([]File, error) {
	var resp struct {
		Data []File `json:"data"`
	}

	req := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp)

	if params != nil {
		if len(params.Fields) > 0 {
			req.SetQueryParam("fields", joinFields(params.Fields))
		}
		if params.Filter != nil {
			req.SetQueryParam("filter", toJSONString(params.Filter))
		}
		if params.Limit > 0 {
			req.SetQueryParam("limit", fmt.Sprintf("%d", params.Limit))
		}
	}

	response, err := req.Get("/files")
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to list files: %s", response.Status())
	}

	return resp.Data, nil
}

// Upload uploads a file from local path
func (s *FilesService) Upload(ctx context.Context, filePath string, metadata map[string]interface{}) (*File, error) {
	var resp struct {
		Data File `json:"data"`
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Convert metadata from map[string]interface{} to map[string]string
	stringMetadata := make(map[string]string)
	for k, v := range metadata {
		stringMetadata[k] = fmt.Sprintf("%v", v)
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetFile("file", filePath).
		SetFormData(stringMetadata).
		SetResult(&resp).
		Post("/files")

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to upload file: %s", response.Status())
	}

	return &resp.Data, nil
}

// Update updates file metadata
func (s *FilesService) Update(ctx context.Context, id string, metadata map[string]interface{}) (*File, error) {
	var resp struct {
		Data File `json:"data"`
	}

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(metadata).
		SetResult(&resp).
		Patch(fmt.Sprintf("/files/%s", id))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to update file: %s", response.Status())
	}

	return &resp.Data, nil
}

// Delete deletes a file
func (s *FilesService) Delete(ctx context.Context, id string) error {
	response, err := s.client.httpClient.R().
		SetContext(ctx).
		Delete(fmt.Sprintf("/files/%s", id))

	if err != nil {
		return err
	}

	if response.StatusCode() != 204 {
		return fmt.Errorf("failed to delete file: %s", response.Status())
	}

	return nil
}
