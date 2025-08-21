package directus

import (
	"encoding/json"
	"time"
)

// Field represents a field in Directus
type Field struct {
	ID         string       `json:"id,omitempty"`
	Collection string       `json:"collection"`
	Field      string       `json:"field"`
	Type       string       `json:"type"`
	Schema     *FieldSchema `json:"schema,omitempty"`
	Meta       *FieldMeta   `json:"meta,omitempty"`
}

// FieldSchema represents the schema of a field
type FieldSchema struct {
	DataType         string `json:"data_type"`
	IsNullable       bool   `json:"is_nullable"`
	IsPrimaryKey     bool   `json:"is_primary_key"`
	MaxLength        *int   `json:"max_length,omitempty"`
	NumericPrecision *int   `json:"numeric_precision,omitempty"`
	NumericScale     *int   `json:"numeric_scale,omitempty"`
}

// FieldMeta represents the meta information of a field
type FieldMeta struct {
	ID             string                 `json:"id,omitempty"`
	Collection     string                 `json:"collection"`
	Field          string                 `json:"field"`
	Special        []string               `json:"special,omitempty"`
	Interface      *string                `json:"interface,omitempty"`
	Options        map[string]interface{} `json:"options,omitempty"`
	Display        *string                `json:"display,omitempty"`
	DisplayOptions map[string]interface{} `json:"display_options,omitempty"`
	Readonly       bool                   `json:"readonly"`
	Hidden         bool                   `json:"hidden"`
	Sort           *int                   `json:"sort,omitempty"`
	Width          *string                `json:"width,omitempty"`
	Group          *string                `json:"group,omitempty"`
	Translations   []Translation          `json:"translations,omitempty"`
	Note           *string                `json:"note,omitempty"`
}

// Collection represents a collection in Directus
type Collection struct {
	Collection string            `json:"collection"`
	Meta       *CollectionMeta   `json:"meta,omitempty"`
	Schema     *CollectionSchema `json:"schema,omitempty"`
	Fields     []Field           `json:"fields,omitempty"`
}

// CollectionMeta represents the meta information of a collection
type CollectionMeta struct {
	ID                    string        `json:"id,omitempty"`
	Collection            string        `json:"collection"`
	Icon                  *string       `json:"icon,omitempty"`
	Note                  *string       `json:"note,omitempty"`
	DisplayTemplate       *string       `json:"display_template,omitempty"`
	Hidden                bool          `json:"hidden"`
	Singleton             bool          `json:"singleton"`
	Translations          []Translation `json:"translations,omitempty"`
	ArchiveField          *string       `json:"archive_field,omitempty"`
	ArchiveAppFilter      bool          `json:"archive_app_filter"`
	ArchiveValue          *string       `json:"archive_value,omitempty"`
	UnarchiveValue        *string       `json:"unarchive_value,omitempty"`
	SortField             *string       `json:"sort_field,omitempty"`
	Accountability        *string       `json:"accountability,omitempty"`
	Color                 *string       `json:"color,omitempty"`
	ItemDuplicationFields []string      `json:"item_duplication_fields,omitempty"`
	Sort                  *int          `json:"sort,omitempty"`
	Group                 *string       `json:"group,omitempty"`
	Collapse              *string       `json:"collapse,omitempty"`
}

// CollectionSchema represents the schema of a collection
type CollectionSchema struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

// Translation represents a translation object
type Translation struct {
	Language    string      `json:"language"`
	Translation interface{} `json:"translation"`
}

// Item represents a generic item in Directus
type Item map[string]interface{}

// File represents a file in Directus
type File struct {
	ID               string          `json:"id,omitempty"`
	Storage          string          `json:"storage"`
	FilenameDisk     string          `json:"filename_disk"`
	FilenameDownload string          `json:"filename_download"`
	Title            *string         `json:"title,omitempty"`
	Type             *string         `json:"type,omitempty"`
	Folder           *string         `json:"folder,omitempty"`
	UploadedBy       *string         `json:"uploaded_by,omitempty"`
	UploadedOn       time.Time       `json:"uploaded_on"`
	ModifiedBy       *string         `json:"modified_by,omitempty"`
	ModifiedOn       time.Time       `json:"modified_on"`
	Filesize         int64           `json:"filesize"`
	Width            *int            `json:"width,omitempty"`
	Height           *int            `json:"height,omitempty"`
	FocalPoint       *string         `json:"focal_point,omitempty"`
	Duration         *int            `json:"duration,omitempty"`
	Description      *string         `json:"description,omitempty"`
	Location         *string         `json:"location,omitempty"`
	Tags             []string        `json:"tags,omitempty"`
	Metadata         json.RawMessage `json:"metadata,omitempty"`
}

// User represents a user in Directus
type User struct {
	ID                 string          `json:"id,omitempty"`
	FirstName          *string         `json:"first_name,omitempty"`
	LastName           *string         `json:"last_name,omitempty"`
	Email              string          `json:"email"`
	Password           *string         `json:"password,omitempty"`
	Location           *string         `json:"location,omitempty"`
	Title              *string         `json:"title,omitempty"`
	Description        *string         `json:"description,omitempty"`
	Tags               []string        `json:"tags,omitempty"`
	Avatar             *string         `json:"avatar,omitempty"`
	Language           *string         `json:"language,omitempty"`
	Theme              *string         `json:"theme,omitempty"`
	TFA                bool            `json:"tfa_secret"`
	Status             string          `json:"status"`
	Role               *string         `json:"role,omitempty"`
	Token              *string         `json:"token,omitempty"`
	LastAccess         time.Time       `json:"last_access"`
	LastPage           *string         `json:"last_page,omitempty"`
	Provider           string          `json:"provider"`
	ExternalIdentifier *string         `json:"external_identifier,omitempty"`
	AuthData           json.RawMessage `json:"auth_data,omitempty"`
}

// Role represents a role in Directus
type Role struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name"`
	Icon        string       `json:"icon"`
	Description *string      `json:"description,omitempty"`
	IPAccess    []string     `json:"ip_access,omitempty"`
	EnforceTFA  bool         `json:"enforce_tfa"`
	AdminAccess bool         `json:"admin_access"`
	AppAccess   bool         `json:"app_access"`
	Permissions []Permission `json:"permissions,omitempty"`
	Users       []User       `json:"users,omitempty"`
}

// Permission represents a permission in Directus
type Permission struct {
	ID          string                 `json:"id,omitempty"`
	Role        string                 `json:"role"`
	Collection  string                 `json:"collection"`
	Action      string                 `json:"action"`
	Permissions map[string]interface{} `json:"permissions,omitempty"`
	Validation  map[string]interface{} `json:"validation,omitempty"`
	Presets     map[string]interface{} `json:"presets,omitempty"`
	Fields      []string               `json:"fields,omitempty"`
}

// QueryParams represents query parameters for filtering and pagination
type QueryParams struct {
	Fields []string               `json:"fields,omitempty"`
	Filter map[string]interface{} `json:"filter,omitempty"`
	Search string                 `json:"search,omitempty"`
	Sort   []string               `json:"sort,omitempty"`
	Limit  int                    `json:"limit,omitempty"`
	Offset int                    `json:"offset,omitempty"`
	Page   int                    `json:"page,omitempty"`
	Deep   map[string]interface{} `json:"deep,omitempty"`
	Export string                 `json:"export,omitempty"`
}

// Response represents a generic API response
type Response struct {
	Data interface{} `json:"data"`
	Meta *Meta       `json:"meta,omitempty"`
}

// Meta represents pagination metadata
type Meta struct {
	FilterCount int `json:"filter_count"`
	TotalCount  int `json:"total_count"`
}

// ErrorResponse represents an error response from the API
type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

// ErrorDetail represents an individual error
type ErrorDetail struct {
	Message    string                 `json:"message"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

// Webhook represents a webhook configuration
type Webhook struct {
	ID          string            `json:"id,omitempty"`
	Name        string            `json:"name"`
	Method      string            `json:"method"`
	URL         string            `json:"url"`
	Status      string            `json:"status"`
	Data        bool              `json:"data"`
	Actions     []string          `json:"actions"`
	Collections []string          `json:"collections"`
	Headers     map[string]string `json:"headers,omitempty"`
}

// Flow represents a flow configuration
type Flow struct {
	ID             string                 `json:"id,omitempty"`
	Name           string                 `json:"name"`
	Icon           string                 `json:"icon"`
	Color          string                 `json:"color"`
	Description    *string                `json:"description,omitempty"`
	Status         string                 `json:"status"`
	Trigger        string                 `json:"trigger"`
	Accountability *string                `json:"accountability,omitempty"`
	Options        map[string]interface{} `json:"options,omitempty"`
	Operation      *FlowOperation         `json:"operation,omitempty"`
}

// FlowOperation represents a flow operation
type FlowOperation struct {
	ID        string                 `json:"id,omitempty"`
	Name      string                 `json:"name"`
	Key       string                 `json:"key"`
	Type      string                 `json:"type"`
	PositionX int                    `json:"position_x"`
	PositionY int                    `json:"position_y"`
	Options   map[string]interface{} `json:"options,omitempty"`
	Resolve   *string                `json:"resolve,omitempty"`
	Reject    *string                `json:"reject,omitempty"`
}
