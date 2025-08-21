package directus

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// joinFields joins field names with commas
func joinFields(fields []string) string {
	return strings.Join(fields, ",")
}

// toJSONString converts a map to JSON string
func toJSONString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return "{}"
	}
	return string(b)
}

// parseError parses an error response from the API
func parseError(resp *resty.Response) error {
	var errResp ErrorResponse
	if err := json.Unmarshal(resp.Body(), &errResp); err != nil {
		return fmt.Errorf("API request failed with status %d: %s", resp.StatusCode(), resp.String())
	}

	if len(errResp.Errors) > 0 {
		return fmt.Errorf("API error: %s", errResp.Errors[0].Message)
	}

	return fmt.Errorf("API request failed with status %d", resp.StatusCode())
}

// isSuccessStatus checks if the status code indicates success
func isSuccessStatus(code int) bool {
	return code >= 200 && code < 300
}

// buildQueryParams builds query parameters from QueryParams
func buildQueryParams(params *QueryParams) map[string]string {
	query := make(map[string]string)
	if params == nil {
		return query
	}

	if len(params.Fields) > 0 {
		query["fields"] = joinFields(params.Fields)
	}
	if params.Filter != nil {
		query["filter"] = toJSONString(params.Filter)
	}
	if params.Search != "" {
		query["search"] = params.Search
	}
	if len(params.Sort) > 0 {
		query["sort"] = joinFields(params.Sort)
	}
	if params.Limit > 0 {
		query["limit"] = fmt.Sprintf("%d", params.Limit)
	}
	if params.Offset > 0 {
		query["offset"] = fmt.Sprintf("%d", params.Offset)
	}
	if params.Page > 0 {
		query["page"] = fmt.Sprintf("%d", params.Page)
	}
	if params.Deep != nil {
		query["deep"] = toJSONString(params.Deep)
	}

	return query
}
