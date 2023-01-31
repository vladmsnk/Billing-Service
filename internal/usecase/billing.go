package usecase

import (
	"context"
	"vladmsnk/billing/internal/dto"
)

// BillingUseCase -.
type BillingUseCase struct {
	billingRepo BillingRepo
}

// NewBillingUseCase -.
func NewBillingUseCase(br BillingRepo) *BillingUseCase {
	return &BillingUseCase{
		billingRepo: br,
	}
}

// MakePurchase -.
func (uc *BillingUseCase) MakePurchase(ctx context.Context, request dto.MakePurchaseRequest) error {
	return nil
}

// GetBalance -
func (uc *BillingUseCase) GetBalance(ctx context.Context) error {
	//getUserCreds
	return nil
}
