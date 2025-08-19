package utils

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/rhyoharianja/go-directusSDK/models"
)

// QueryBuilder helps build query parameters
type QueryBuilder struct {
	params *models.QueryParams
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		params: &models.QueryParams{},
	}
}

// Fields sets the fields to retrieve
func (qb *QueryBuilder) Fields(fields ...string) *QueryBuilder {
	qb.params.Fields = fields
	return qb
}

// Filter sets the filter criteria
func (qb *QueryBuilder) Filter(filter map[string]interface{}) *QueryBuilder {
	qb.params.Filter = filter
	return qb
}

// Search sets the search query
func (qb *QueryBuilder) Search(search string) *QueryBuilder {
	qb.params.Search = search
	return qb
}

// Sort sets the sort order
func (qb *QueryBuilder) Sort(sort ...string) *QueryBuilder {
	qb.params.Sort = sort
	return qb
}

// Limit sets the limit
func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.params.Limit = limit
	return qb
}

// Offset sets the offset
func (qb *QueryBuilder) Offset(offset int) *QueryBuilder {
	qb.params.Offset = offset
	return qb
}

// Page sets the page number
func (qb *QueryBuilder) Page(page int) *QueryBuilder {
	qb.params.Page = page
	return qb
}

// Build converts the query builder to QueryParams
func (qb *QueryBuilder) Build() *models.QueryParams {
	return qb.params
}

// ToURLValues converts QueryParams to url.Values
func ToURLValues(params *models.QueryParams) url.Values {
	values := url.Values{}
	
	if len(params.Fields) > 0 {
		values.Set("fields", strings.Join(params.Fields, ","))
	}
	
	if params.Limit > 0 {
		values.Set("limit", strconv.Itoa(params.Limit))
	}
	
	if params.Offset > 0 {
		values.Set("offset", strconv.Itoa(params.Offset))
	}
	
	if params.Page > 0 {
		values.Set("page", strconv.Itoa(params.Page))
	}
	
	if params.Search != "" {
		values.Set("search", params.Search)
	}
	
	if len(params.Sort) > 0 {
		values.Set("sort", strings.Join(params.Sort, ","))
	}
	
	return values
}
