package directus

import (
	"context"
	"fmt"
	"net/http"
)

// ItemsService handles operations on collection items
type ItemsService struct {
	client *Client
}

// NewItemsService creates a new items service
func NewItemsService(client *Client) *ItemsService {
	return &ItemsService{client: client}
}

// Get retrieves a single item by ID
func (s *ItemsService) Get(ctx context.Context, collection string, id string, params *QueryParams) (Item, error) {
	var resp Response
	path := fmt.Sprintf("/items/%s/%s", collection, id)

	req := s.client.httpClient.R().
		SetContext(ctx).
		SetResult(&resp)

	if params != nil {
		if len(params.Fields) > 0 {
			req.SetQueryParam("fields", joinFields(params.Fields))
		}
		if params.Deep != nil {
			req.SetQueryParam("deep", toJSONString(params.Deep))
		}
	}

	response, err := req.Get(path)
	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, parseError(response)
	}

	if resp.Data == nil {
		return nil, fmt.Errorf("no data returned")
	}

	item, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	return Item(item), nil
}

// List retrieves multiple items from a collection
func (s *ItemsService) List(ctx context.Context, collection string, params *QueryParams) ([]Item, *Meta, error) {
	var resp Response
	path := fmt.Sprintf("/items/%s", collection)

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
		if params.Search != "" {
			req.SetQueryParam("search", params.Search)
		}
		if len(params.Sort) > 0 {
			req.SetQueryParam("sort", joinFields(params.Sort))
		}
		if params.Limit > 0 {
			req.SetQueryParam("limit", fmt.Sprintf("%d", params.Limit))
		}
		if params.Offset > 0 {
			req.SetQueryParam("offset", fmt.Sprintf("%d", params.Offset))
		}
		if params.Page > 0 {
			req.SetQueryParam("page", fmt.Sprintf("%d", params.Page))
		}
		if params.Deep != nil {
			req.SetQueryParam("deep", toJSONString(params.Deep))
		}
	}

	response, err := req.Get(path)
	if err != nil {
		return nil, nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, nil, parseError(response)
	}

	data, ok := resp.Data.([]interface{})
	if !ok {
		return nil, nil, fmt.Errorf("invalid response format")
	}

	items := make([]Item, len(data))
	for i, v := range data {
		if item, ok := v.(map[string]interface{}); ok {
			items[i] = Item(item)
		}
	}

	return items, resp.Meta, nil
}

// Create creates a new item in a collection
func (s *ItemsService) Create(ctx context.Context, collection string, item Item) (Item, error) {
	var resp Response
	path := fmt.Sprintf("/items/%s", collection)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(item).
		SetResult(&resp).
		Post(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusCreated {
		return nil, parseError(response)
	}

	createdItem, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	return Item(createdItem), nil
}

// Update updates an existing item in a collection
func (s *ItemsService) Update(ctx context.Context, collection string, id string, item Item) (Item, error) {
	var resp Response
	path := fmt.Sprintf("/items/%s/%s", collection, id)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(item).
		SetResult(&resp).
		Patch(path)

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != http.StatusOK {
		return nil, parseError(response)
	}

	updatedItem, ok := resp.Data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("invalid response format")
	}

	return Item(updatedItem), nil
}

// Delete deletes an item from a collection
func (s *ItemsService) Delete(ctx context.Context, collection string, id string) error {
	path := fmt.Sprintf("/items/%s/%s", collection, id)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		Delete(path)

	if err != nil {
		return err
	}

	if response.StatusCode() != http.StatusNoContent {
		return parseError(response)
	}

	return nil
}

// DeleteMultiple deletes multiple items from a collection
func (s *ItemsService) DeleteMultiple(ctx context.Context, collection string, ids []string) error {
	path := fmt.Sprintf("/items/%s", collection)

	response, err := s.client.httpClient.R().
		SetContext(ctx).
		SetBody(map[string][]string{"keys": ids}).
		Delete(path)

	if err != nil {
		return err
	}

	if response.StatusCode() != http.StatusOK {
		return parseError(response)
	}

	return nil
}
