package repositories

import (
	"context"
	"fmt"

	"github.com/rhyoharianja/go-directusSDK/client"
	"github.com/rhyoharianja/go-directusSDK/models"
)

type authRepository struct {
	client *client.HTTPClient
}

// NewAuthRepository creates a new auth repository
func NewAuthRepository(client *client.HTTPClient) AuthRepository {
	return &authRepository{
		client: client,
	}
}

func (r *authRepository) Login(ctx context.Context, email, password string) (*models.AuthResponse, error) {
	req := &models.LoginRequest{
		Email:    email,
		Password: password,
		Mode:     "json",
	}

	var resp models.AuthResponse
	if err := r.client.Post(ctx, "/auth/login", req, &resp); err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	return &resp, nil
}

func (r *authRepository) Logout(ctx context.Context) error {
	if err := r.client.Post(ctx, "/auth/logout", nil, nil); err != nil {
		return fmt.Errorf("logout failed: %w", err)
	}
	return nil
}

func (r *authRepository) Refresh(ctx context.Context, refreshToken string) (*models.AuthResponse, error) {
	req := &models.RefreshTokenRequest{
		RefreshToken: refreshToken,
	}

	var resp models.AuthResponse
	if err := r.client.Post(ctx, "/auth/refresh", req, &resp); err != nil {
		return nil, fmt.Errorf("refresh failed: %w", err)
	}

	return &resp, nil
}

func (r *authRepository) Me(ctx context.Context) (*models.User, error) {
	var user models.User
	if err := r.client.Get(ctx, "/users/me", nil, &user); err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}
	return &user, nil
}
