package models

import (
	"time"
)

// BaseModel contains common fields for all Directus models
type BaseModel struct {
	ID        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// User represents a Directus user
type User struct {
	BaseModel
	Email         string                 `json:"email"`
	FirstName     string                 `json:"first_name"`
	LastName      string                 `json:"last_name"`
	Avatar        *string                `json:"avatar,omitempty"`
	Role          *string                `json:"role,omitempty"`
	Token         *string                `json:"token,omitempty"`
	LastAccess    *time.Time             `json:"last_access,omitempty"`
	Provider      string                 `json:"provider,omitempty"`
	ExternalID    *string                `json:"external_identifier,omitempty"`
	AuthData      map[string]interface{} `json:"auth_data,omitempty"`
	EmailVerified bool                   `json:"email_verified,omitempty"`
	Theme         string                 `json:"theme,omitempty"`
	Language      string                 `json:"language,omitempty"`
	Status        string                 `json:"status,omitempty"`
}

// Role represents a Directus role
type Role struct {
	BaseModel
	Name        string                 `json:"name"`
	Icon        string                 `json:"icon,omitempty"`
	Description string                 `json:"description,omitempty"`
	IPAccess    []string               `json:"ip_access,omitempty"`
	Enforce2FA  bool                   `json:"enforce_tfa,omitempty"`
	AdminAccess bool                   `json:"admin_access,omitempty"`
	AppAccess   bool                   `json:"app_access,omitempty"`
	Users       []User                 `json:"users,omitempty"`
	Collections []string               `json:"collections,omitempty"`
	Permissions []Permission           `json:"permissions,omitempty"`
}

// Permission represents a Directus permission
type Permission struct {
	BaseModel
	Role          string                 `json:"role"`
	Collection    string                 `json:"collection"`
	Action        string                 `json:"action"`
	Permissions   map[string]interface{} `json:"permissions,omitempty"`
	Validation    map[string]interface{} `json:"validation,omitempty"`
	Presets       map[string]interface{} `json:"presets,omitempty"`
	Fields        []string               `json:"fields,omitempty"`
	Limit         *int                   `json:"limit,omitempty"`
	System        bool                   `json:"system,omitempty"`
}

// File represents a Directus file
type File struct {
	BaseModel
	Storage        string                 `json:"storage"`
	FilenameDisk   string                 `json:"filename_disk"`
	FilenameDownload string               `json:"filename_download"`
	Title          string                 `json:"title"`
	Type           string                 `json:"type"`
	Folder         *string                `json:"folder,omitempty"`
	UploadedBy     *string                `json:"uploaded_by,omitempty"`
	UploadedOn     time.Time              `json:"uploaded_on"`
	ModifiedBy     *string                `json:"modified_by,omitempty"`
	ModifiedOn     time.Time              `json:"modified_on"`
	Filesize       int64                  `json:"filesize"`
	Width          *int                   `json:"width,omitempty"`
	Height         *int                   `json:"height,omitempty"`
	FocalPointX    *float64               `json:"focal_point_x,omitempty"`
	FocalPointY    *float64               `json:"focal_point_y,omitempty"`
	Duration       *int                   `json:"duration,omitempty"`
	Description    *string                `json:"description,omitempty"`
	Location       *string                `json:"location,omitempty"`
	Tags           []string               `json:"tags,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}

// Collection represents a Directus collection
type Collection struct {
	BaseModel
	Collection  string                 `json:"collection"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
	Schema      map[string]interface{} `json:"schema,omitempty"`
	Fields      []Field                `json:"fields,omitempty"`
}

// Field represents a field in a Directus collection
type Field struct {
	BaseModel
	Collection string                 `json:"collection"`
	Field      string                 `json:"field"`
	Type       string                 `json:"type"`
	Meta       map[string]interface{} `json:"meta,omitempty"`
	Schema     map[string]interface{} `json:"schema,omitempty"`
}

// QueryParams represents query parameters for API requests
type QueryParams struct {
	Fields    []string               `json:"fields,omitempty"`
	Filter    map[string]interface{} `json:"filter,omitempty"`
	Search    string                 `json:"search,omitempty"`
	Sort      []string               `json:"sort,omitempty"`
	Limit     int                    `json:"limit,omitempty"`
	Offset    int                    `json:"offset,omitempty"`
	Page      int                    `json:"page,omitempty"`
	Deep      map[string]interface{} `json:"deep,omitempty"`
	Aggregate map[string]interface{} `json:"aggregate,omitempty"`
	GroupBy   []string               `json:"group_by,omitempty"`
}

// Pagination represents pagination information
type Pagination struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"total_pages"`
}

// APIResponse represents a generic API response
type APIResponse struct {
	Data       interface{} `json:"data"`
	Meta       interface{} `json:"meta,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// AuthResponse represents authentication response
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Expires      int64  `json:"expires"`
}

// LoginRequest represents login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Mode     string `json:"mode,omitempty"`
}

// RefreshTokenRequest represents refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// CreateRequest represents create request
type CreateRequest struct {
	Data interface{} `json:"data"`
}

// UpdateRequest represents update request
type UpdateRequest struct {
	Data interface{} `json:"data"`
}

// DeleteRequest represents delete request
type DeleteRequest struct {
	Keys []string `json:"keys"`
}

// FileUploadRequest represents file upload request
type FileUploadRequest struct {
	File        []byte                 `json:"file"`
	Filename    string                 `json:"filename"`
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Folder      string                 `json:"folder,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}
