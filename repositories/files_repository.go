package repositories

import (
	"context"
	"fmt"
	"net/url"

	"github.com/rhyoharianja/go-directusSDK/client"
	"github.com/rhyoharianja/go-directusSDK/models"
)

type filesRepository struct {
	client *client.HTTPClient
}

// NewFilesRepository creates a new files repository
func NewFilesRepository(client *client.HTTPClient) FilesRepository {
	return &filesRepository{
		client: client,
	}
}

func (r *filesRepository) GetMany(ctx context.Context, params *models.QueryParams) ([]*models.File, error) {
	var paramsMap url.Values
	if params != nil {
		paramsMap = url.Values{}
		if len(params.Fields) > 0 {
			for _, field := range params.Fields {
				paramsMap.Add("fields", field)
			}
		}
		if params.Limit > 0 {
			paramsMap.Add("limit", fmt.Sprintf("%d", params.Limit))
		}
		if params.Offset > 0 {
			paramsMap.Add("offset", fmt.Sprintf("%d", params.Offset))
		}
		if len(params.Sort) > 0 {
			for _, sort := range params.Sort {
				paramsMap.Add("sort", sort)
			}
		}
	}

	var response struct {
		Data []*models.File `json:"data"`
	}

	if err := r.client.Get(ctx, "/files", paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get files failed: %w", err)
	}

	return response.Data, nil
}

func (r *filesRepository) GetOne(ctx context.Context, id string, params *models.QueryParams) (*models.File, error) {
	var paramsMap url.Values
	if params != nil {
		paramsMap = url.Values{}
		if len(params.Fields) > 0 {
			for _, field := range params.Fields {
				paramsMap.Add("fields", field)
			}
		}
	}

	var response struct {
		Data *models.File `json:"data"`
	}

	path := fmt.Sprintf("/files/%s", id)
	if err := r.client.Get(ctx, path, paramsMap, &response); err != nil {
		return nil, fmt.Errorf("get file failed: %w", err)
	}

	return response.Data, nil
}

func (r *filesRepository) Upload(ctx context.Context, file *models.FileUploadRequest) (*models.File, error) {
	var response struct {
		Data *models.File `json:"data"`
	}

	if err := r.client.Post(ctx, "/files", file, &response); err != nil {
		return nil, fmt.Errorf("upload file failed: %w", err)
	}

	return response.Data, nil
}

func (r *filesRepository) Update(ctx context.Context, id string, data interface{}) (*models.File, error) {
	var response struct {
		Data *models.File `json:"data"`
	}

	path := fmt.Sprintf("/files/%s", id)
	if err := r.client.Patch(ctx, path, data, &response); err != nil {
		return nil, fmt.Errorf("update file failed: %w", err)
	}

	return response.Data, nil
}

func (r *filesRepository) Delete(ctx context.Context, id string) error {
	path := fmt.Sprintf("/files/%s", id)
	if err := r.client.Delete(ctx, path); err != nil {
		return fmt.Errorf("delete file failed: %w", err)
	}

	return nil
}
