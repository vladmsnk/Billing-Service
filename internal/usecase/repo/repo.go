package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"vladmsnk/billing/internal/entity"
	"vladmsnk/billing/pkg/postgres"
)

// Repo -.
type Repo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *Repo {
	return &Repo{pg}
}

// CreateUser -.
func (r *Repo) CreateUser(ctx context.Context, user entity.User) error {

	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx, InsertUser, user.ID, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	tx.Commit(ctx)
	return nil
}

func (r *Repo) CheckUserExists(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	tx, err := r.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	err = tx.QueryRow(ctx, CheckUserExists, email).Scan(&user)
	if err != nil {
		return nil, err
	}
	tx.Commit(ctx)
	return &user, nil
}

func (r *Repo) GetBalance(ctx context.Context, userID uuid.UUID) {

}
