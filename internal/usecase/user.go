package usecase

import (
	"context"
	"github.com/google/uuid"
	"vladmsnk/billing/internal/dto"
	"vladmsnk/billing/internal/jwt"
)

// AuthUseCase -.
type AuthUseCase struct {
	authRepo AuthRepo
}

// NewAuthUseCase -.
func NewAuthUseCase(ar AuthRepo) *AuthUseCase {
	return &AuthUseCase{
		authRepo: ar,
	}
}

// CreateUser -.
func (uc *AuthUseCase) CreateUser(ctx context.Context, request dto.User) error {
	userEmail := request.Email
	_, err := uc.authRepo.CheckUserExists(ctx, userEmail)
	if err != nil {
		return err
	}

	userEntity := request.FromDto(uuid.New())
	//todo плохо its bad to hash password on backend side
	err = userEntity.HashPassword(request.Password)
	if err != nil {
		return err
	}

	err = uc.authRepo.CreateUser(ctx, userEntity)
	if err != nil {
		return err
	}

	return nil
}

// GenerateToken -.
func (uc *AuthUseCase) GenerateToken(ctx context.Context, request dto.TokenRequest) (string, error) {
	user, err := uc.authRepo.CheckUserExists(ctx, request.Email)
	if err != nil {
		return "", err
	}

	credErr := user.CheckPassword(request.Password)
	if credErr != nil {
		return "", credErr
	}

	tokenString, err := jwt.GenerateJWT(user.Email, user.Username)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
