package utils

import (
	"fmt"
	"strings"
)

// QueryAlias provides a fluent interface for building complex Directus queries
type QueryAlias struct {
	fields    []string
	filter    map[string]interface{}
	search    string
	sort      []string
	limit     int
	offset    int
	page      int
	deep      map[string]interface{}
	aggregate map[string]interface{}
	groupBy   []string
}

// New creates a new QueryAlias instance
func New() *QueryAlias {
	return &QueryAlias{
		fields:    []string{},
		filter:    make(map[string]interface{}),
		sort:      []string{},
		deep:      make(map[string]interface{}),
		aggregate: make(map[string]interface{}),
		groupBy:   []string{},
	}
}

// Select adds fields to select
func (q *QueryAlias) Select(fields ...string) *QueryAlias {
	q.fields = append(q.fields, fields...)
	return q
}

// Where adds a filter condition
func (q *QueryAlias) Where(field string, operator string, value interface{}) *QueryAlias {
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	q.filter[field] = map[string]interface{}{
		operator: value,
	}
	return q
}

// WhereIn adds an IN filter condition
func (q *QueryAlias) WhereIn(field string, values []interface{}) *QueryAlias {
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	q.filter[field] = map[string]interface{}{
		"_in": values,
	}
	return q
}

// WhereNotIn adds a NOT IN filter condition
func (q *QueryAlias) WhereNotIn(field string, values []interface{}) *QueryAlias {
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	q.filter[field] = map[string]interface{}{
		"_nin": values,
	}
	return q
}

// WhereBetween adds a BETWEEN filter condition
func (q *QueryAlias) WhereBetween(field string, min, max interface{}) *QueryAlias {
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	q.filter[field] = map[string]interface{}{
		"_between": []interface{}{min, max},
	}
	return q
}

// WhereNull adds an IS NULL filter condition
func (q *QueryAlias) WhereNull(field string) *QueryAlias {
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	q.filter[field] = map[string]interface{}{
		"_null": true,
	}
	return q
}

// WhereNotNull adds an IS NOT NULL filter condition
func (q *QueryAlias) WhereNotNull(field string) *QueryAlias {
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	q.filter[field] = map[string]interface{}{
		"_nnull": true,
	}
	return q
}

// WhereLike adds a LIKE filter condition
func (q *QueryAlias) WhereLike(field string, pattern string) *QueryAlias {
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	q.filter[field] = map[string]interface{}{
		"_like": pattern,
	}
	return q
}

// WhereContains adds a CONTAINS filter condition
func (q *QueryAlias) WhereContains(field string, value interface{}) *QueryAlias {
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	q.filter[field] = map[string]interface{}{
		"_contains": value,
	}
	return q
}

// Search adds a search query
func (q *QueryAlias) Search(query string) *QueryAlias {
	q.search = query
	return q
}

// OrderBy adds sorting
func (q *QueryAlias) OrderBy(field string, direction string) *QueryAlias {
	direction = strings.ToUpper(direction)
	if direction != "ASC" && direction != "DESC" {
		direction = "ASC"
	}
	q.sort = append(q.sort, fmt.Sprintf("%s:%s", field, direction))
	return q
}

// OrderByDesc adds descending sort
func (q *QueryAlias) OrderByDesc(field string) *QueryAlias {
	q.sort = append(q.sort, fmt.Sprintf("-%s", field))
	return q
}

// OrderByAsc adds ascending sort
func (q *QueryAlias) OrderByAsc(field string) *QueryAlias {
	q.sort = append(q.sort, field)
	return q
}

// Limit sets the limit
func (q *QueryAlias) Limit(limit int) *QueryAlias {
	q.limit = limit
	return q
}

// Offset sets the offset
func (q *QueryAlias) Offset(offset int) *QueryAlias {
	q.offset = offset
	return q
}

// Page sets the page number
func (q *QueryAlias) Page(page int) *QueryAlias {
	q.page = page
	return q
}

// Deep adds deep querying
func (q *QueryAlias) Deep(relation string, query *QueryAlias) *QueryAlias {
	if q.deep == nil {
		q.deep = make(map[string]interface{})
	}
	q.deep[relation] = query.Build()
	return q
}

// Aggregate adds aggregation
func (q *QueryAlias) Aggregate(function string, field string) *QueryAlias {
	if q.aggregate == nil {
		q.aggregate = make(map[string]interface{})
	}
	q.aggregate[function] = field
	return q
}

// GroupBy adds grouping
func (q *QueryAlias) GroupBy(fields ...string) *QueryAlias {
	q.groupBy = append(q.groupBy, fields...)
	return q
}

// Build converts the query alias to a map
func (q *QueryAlias) Build() map[string]interface{} {
	result := make(map[string]interface{})
	
	if len(q.fields) > 0 {
		result["fields"] = q.fields
	}
	
	if len(q.filter) > 0 {
		result["filter"] = q.filter
	}
	
	if q.search != "" {
		result["search"] = q.search
	}
	
	if len(q.sort) > 0 {
		result["sort"] = q.sort
	}
	
	if q.limit > 0 {
		result["limit"] = q.limit
	}
	
	if q.offset > 0 {
		result["offset"] = q.offset
	}
	
	if q.page > 0 {
		result["page"] = q.page
	}
	
	if len(q.deep) > 0 {
		result["deep"] = q.deep
	}
	
	if len(q.aggregate) > 0 {
		result["aggregate"] = q.aggregate
	}
	
	if len(q.groupBy) > 0 {
		result["group_by"] = q.groupBy
	}
	
	return result
}

// And adds an AND condition
func (q *QueryAlias) And(conditions ...*QueryAlias) *QueryAlias {
	if len(conditions) == 0 {
		return q
	}
	
	if q.filter == nil {
		q.filter = make(map[string]interface{})
	}
	
	andConditions := make([]map[string]interface{}, 0)
	
	// Add existing filter as first condition if not empty
	if len(q.filter) > 0 {
		andConditions = append(andConditions, q.filter)
	}
	
	// Add new conditions
	for _, condition := range conditions {
		if len(condition.filter) > 0 {
			andConditions = append(andConditions, condition.filter)
		}
	}
	
	if len(andConditions) > 0 {
		q.filter = map[string]interface{}{
			"_and": andConditions,
		}
	}
	
	return q
}

// Or adds an OR condition
func (q *QueryAlias) Or(conditions ...*QueryAlias) *QueryAlias {
	if len(conditions) == 0 {
		return q
	}
	
	orConditions := make([]map[string]interface{}, 0)
	
	// Add existing filter as first condition if not empty
	if len(q.filter) > 0 {
		orConditions = append(orConditions, q.filter)
	}
	
	// Add new conditions
	for _, condition := range conditions {
		if len(condition.filter) > 0 {
			orConditions = append(orConditions, condition.filter)
		}
	}
	
	if len(orConditions) > 0 {
		q.filter = map[string]interface{}{
			"_or": orConditions,
		}
	}
	
	return q
}

// Not adds a NOT condition
func (q *QueryAlias) Not(condition *QueryAlias) *QueryAlias {
	if len(condition.filter) > 0 {
		if q.filter == nil {
			q.filter = make(map[string]interface{})
		}
		q.filter["_not"] = condition.filter
	}
	return q
}

// With adds related data
func (q *QueryAlias) With(relation string, query *QueryAlias) *QueryAlias {
	return q.Deep(relation, query)
}

// Count adds count aggregation
func (q *QueryAlias) Count(field string) *QueryAlias {
	return q.Aggregate("count", field)
}

// Sum adds sum aggregation
func (q *QueryAlias) Sum(field string) *QueryAlias {
	return q.Aggregate("sum", field)
}

// Avg adds average aggregation
func (q *QueryAlias) Avg(field string) *QueryAlias {
	return q.Aggregate("avg", field)
}

// Min adds minimum aggregation
func (q *QueryAlias) Min(field string) *QueryAlias {
	return q.Aggregate("min", field)
}

// Max adds maximum aggregation
func (q *QueryAlias) Max(field string) *QueryAlias {
	return q.Aggregate("max", field)
}

// Clone creates a copy of the query alias
func (q *QueryAlias) Clone() *QueryAlias {
	newQuery := New()
	
	if len(q.fields) > 0 {
		newQuery.fields = append([]string{}, q.fields...)
	}
	
	if len(q.filter) > 0 {
		newQuery.filter = make(map[string]interface{})
		for k, v := range q.filter {
			newQuery.filter[k] = v
		}
	}
	
	if q.search != "" {
		newQuery.search = q.search
	}
	
	if len(q.sort) > 0 {
		newQuery.sort = append([]string{}, q.sort...)
	}
	
	if q.limit > 0 {
		newQuery.limit = q.limit
	}
	
	if q.offset > 0 {
		newQuery.offset = q.offset
	}
	
	if q.page > 0 {
		newQuery.page = q.page
	}
	
	if len(q.deep) > 0 {
		newQuery.deep = make(map[string]interface{})
		for k, v := range q.deep {
			newQuery.deep[k] = v
		}
	}
	
	if len(q.aggregate) > 0 {
		newQuery.aggregate = make(map[string]interface{})
		for k, v := range q.aggregate {
			newQuery.aggregate[k] = v
		}
	}
	
	if len(q.groupBy) > 0 {
		newQuery.groupBy = append([]string{}, q.groupBy...)
	}
	
	return newQuery
}

// Reset clears all query parameters
func (q *QueryAlias) Reset() *QueryAlias {
	q.fields = []string{}
	q.filter = make(map[string]interface{})
	q.search = ""
	q.sort = []string{}
	q.limit = 0
	q.offset = 0
	q.page = 0
	q.deep = make(map[string]interface{})
	q.aggregate = make(map[string]interface{})
	q.groupBy = []string{}
	return q
}

// String returns a string representation of the query
func (q *QueryAlias) String() string {
	return fmt.Sprintf("QueryAlias{fields: %v, filter: %v, search: %q, sort: %v, limit: %d, offset: %d, page: %d}",
		q.fields, q.filter, q.search, q.sort, q.limit, q.offset, q.page)
}
