// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"github.com/google/uuid"
	"vladmsnk/billing/internal/dto"
	"vladmsnk/billing/internal/entity"
)

type (
	// Billing -.
	Billing interface {
		MakePurchase(ctx context.Context, request dto.MakePurchaseRequest) error
		GetBalance(ctx context.Context) error
	}

	// Auth -. {
	Auth interface {
		CreateUser(ctx context.Context, request dto.User) error
		GenerateToken(ctx context.Context, request dto.TokenRequest) (string, error)
	}

	// AuthRepo -.
	AuthRepo interface {
		CreateUser(ctx context.Context, user entity.User) error
		CheckUserExists(ctx context.Context, email string) (*entity.User, error)
	}

	// BillingRepo -.
	BillingRepo interface {
		GetBalance(ctx context.Context, userID uuid.UUID)
	}
)
