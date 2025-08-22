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

// parseError parses an error response from the API with improved JSON handling
func parseError(resp *resty.Response) error {
	var errResp ErrorResponse

	// Use safeUnmarshal for better error handling
	if err := safeUnmarshal(resp.Body(), &errResp); err != nil {
		// If JSON parsing fails, try to extract error message from response body
		body := string(resp.Body())
		if body != "" {
			return fmt.Errorf("API request failed with status %d: %s (JSON parse error: %v)",
				resp.StatusCode(), body, err)
		}
		return fmt.Errorf("API request failed with status %d: %v", resp.StatusCode(), err)
	}

	if len(errResp.Errors) > 0 {
		// Include extension details if available
		errorMsg := errResp.Errors[0].Message
		if extensions := errResp.Errors[0].Extensions; extensions != nil {
			if code, ok := extensions["code"].(string); ok {
				errorMsg = fmt.Sprintf("%s (code: %s)", errorMsg, code)
			}
		}
		return fmt.Errorf("API error: %s", errorMsg)
	}

	return fmt.Errorf("API request failed with status %d", resp.StatusCode())
}

// parseResponse safely parses API response with improved error handling
func parseResponse(resp *resty.Response, result interface{}) error {
	if !isSuccessStatus(resp.StatusCode()) {
		return parseError(resp)
	}

	if err := safeUnmarshal(resp.Body(), result); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	return nil
}

// isSuccessStatus checks if the status code indicates success
func isSuccessStatus(code int) bool {
	return code >= 200 && code < 300
}

// NewFilterEqual creates an equality filter condition
func NewFilterEqual(field string, value interface{}) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{string(FilterEqual): value}}
}

// NewFilterNotEqual creates a not-equal filter condition
func NewFilterNotEqual(field string, value interface{}) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{string(FilterNotEqual): value}}
}

// NewFilterContains creates a contains filter condition
func NewFilterContains(field string, value string) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{string(FilterContains): value}}
}

// NewFilterIn creates an "in" filter condition
func NewFilterIn(field string, values []interface{}) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{string(FilterIn): values}}
}

// NewFilterBetween creates a between filter condition
func NewFilterBetween(field string, from, to interface{}) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{string(FilterBetween): []interface{}{from, to}}}
}

// NewFilterNull creates a null filter condition
func NewFilterNull(field string) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{string(FilterNull): true}}
}

// NewFilterNotNull creates a not-null filter condition
func NewFilterNotNull(field string) map[string]interface{} {
	return map[string]interface{}{field: map[string]interface{}{string(FilterNotNull): true}}
}

// NewFilterAnd combines multiple filters with AND logic
func NewFilterAnd(filters ...map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{string(LogicalAnd): filters}
}

// NewFilterOr combines multiple filters with OR logic
func NewFilterOr(filters ...map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{string(LogicalOr): filters}
}

// AddAlias adds a field alias to QueryParams
func (qp *QueryParams) AddAlias(originalField, alias string) {
	if qp.Aliases == nil {
		qp.Aliases = make(map[string]string)
	}
	qp.Aliases[originalField] = alias
}

// SetLanguage sets the language for translations
func (qp *QueryParams) SetLanguage(lang string) {
	qp.Lang = lang
}

// validateJSON validates if a byte slice contains valid JSON
func validateJSON(data []byte) error {
	var js interface{}
	return json.Unmarshal(data, &js)
}

// safeUnmarshal safely unmarshals JSON with better error handling
func safeUnmarshal(data []byte, v interface{}) error {
	if err := validateJSON(data); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}
	return json.Unmarshal(data, v)
}
