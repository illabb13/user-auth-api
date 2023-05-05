package user

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/illabb13/users_auth/internal/model"
)

const tableName = "users"

type Repository interface {
	Create(ctx context.Context, model *model.User) error
	Delete(ctx context.Context, model *model.Query) error
}

type repository struct {
	conPool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		conPool: pool,
	}
}

func (r *repository) Create(ctx context.Context, model *model.User) error {
	builder := sq.Insert(tableName).PlaceholderFormat(sq.Dollar).
		Columns("username", "email", "password", "role").
		Values(model.Username, model.Email, model.Password, model.Role)

	query, values, err := builder.ToSql()
	if err != nil {
		return err
	}

	log.Printf("insert query with params: %s, %s", query, values)
	rows, err := r.conPool.Query(ctx, query, values...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *repository) Delete(ctx context.Context, model *model.Query) error {
	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).
		Where("username = ?", model.Username)

	query, values, err := builder.ToSql()
	if err != nil {
		return err
	}

	log.Printf("delete query with params: %s, %s", query, values)
	rows, err := r.conPool.Query(ctx, query, values...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
