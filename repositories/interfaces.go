package repositories

import (
	"context"

	"github.com/rhyoharianja/go-directusSDK/models"
)

// AuthRepository handles authentication operations
type AuthRepository interface {
	Login(ctx context.Context, email, password string) (*models.AuthResponse, error)
	Logout(ctx context.Context) error
	Refresh(ctx context.Context, refreshToken string) (*models.AuthResponse, error)
	Me(ctx context.Context) (*models.User, error)
}

// ItemsRepository handles CRUD operations for items in collections
type ItemsRepository interface {
	GetMany(ctx context.Context, params *models.QueryParams) ([]map[string]interface{}, error)
	GetOne(ctx context.Context, id string, params *models.QueryParams) (map[string]interface{}, error)
	Create(ctx context.Context, data interface{}) (map[string]interface{}, error)
	Update(ctx context.Context, id string, data interface{}) (map[string]interface{}, error)
	Delete(ctx context.Context, id string) error
	DeleteMany(ctx context.Context, ids []string) error
	WithCollection(collection string) ItemsRepository
}

// FilesRepository handles file operations
type FilesRepository interface {
	GetMany(ctx context.Context, params *models.QueryParams) ([]*models.File, error)
	GetOne(ctx context.Context, id string, params *models.QueryParams) (*models.File, error)
	Upload(ctx context.Context, file *models.FileUploadRequest) (*models.File, error)
	Update(ctx context.Context, id string, data interface{}) (*models.File, error)
	Delete(ctx context.Context, id string) error
}

// UsersRepository handles user operations
type UsersRepository interface {
	GetMany(ctx context.Context, params *models.QueryParams) ([]*models.User, error)
	GetOne(ctx context.Context, id string, params *models.QueryParams) (*models.User, error)
	Create(ctx context.Context, user *models.User) (*models.User, error)
	Update(ctx context.Context, id string, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
}

// RolesRepository handles role operations
type RolesRepository interface {
	GetMany(ctx context.Context, params *models.QueryParams) ([]*models.Role, error)
	GetOne(ctx context.Context, id string, params *models.QueryParams) (*models.Role, error)
	Create(ctx context.Context, role *models.Role) (*models.Role, error)
	Update(ctx context.Context, id string, role *models.Role) (*models.Role, error)
	Delete(ctx context.Context, id string) error
}

// CollectionsRepository handles collection operations
type CollectionsRepository interface {
	GetMany(ctx context.Context, params *models.QueryParams) ([]*models.Collection, error)
	GetOne(ctx context.Context, name string, params *models.QueryParams) (*models.Collection, error)
	Create(ctx context.Context, collection *models.Collection) (*models.Collection, error)
	Update(ctx context.Context, name string, collection *models.Collection) (*models.Collection, error)
	Delete(ctx context.Context, name string) error
}
