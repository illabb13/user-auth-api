package user

import (
	"context"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"

	model "github.com/illabb13/users_auth/internal/model/user"
)

const tableName = "users"

type Repository interface {
	Create(ctx context.Context, model *model.UserNewData) error
	Get(ctx context.Context, query *model.Query) (*model.UserInfo, error)
	Delete(ctx context.Context, query *model.Query) error
	Update(ctx context.Context, query *model.Query, user *model.UserNewData) error
}

type repository struct {
	conPool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *repository {
	return &repository{
		conPool: pool,
	}
}

func (r *repository) Create(ctx context.Context, user *model.UserNewData) error {
	builder := sq.Insert(tableName).PlaceholderFormat(sq.Dollar).
		Columns("username", "email", "password", "role").
		Values(user.Username, user.Email, user.Password, user.Role)

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

func (r *repository) Get(ctx context.Context, query *model.Query) (*model.UserInfo, error) {
	builder := sq.Select("username", "email", "role", "created_at", "updated_at").
		From(tableName).
		Where("username = ?", query.Username).
		PlaceholderFormat(sq.Dollar)
	sql, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	log.Printf("select query with params: %s, %s", sql, args)
	rows, err := r.conPool.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user model.UserInfo
	row := rows.Next()
	if row {
		err = rows.Scan(&user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (r *repository) Update(ctx context.Context, query *model.Query, user *model.UserNewData) error {
	builder := sq.Update(tableName).PlaceholderFormat(sq.Dollar).
		Set("username", user.Username).
		Set("email", user.Email).
		Set("password", user.Password).
		Set("role", user.Role).
		Set("updated_at", sq.Expr("NOW()")).
		Where("username = ?", query.Username)

	sql, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	log.Printf("update query with params: %s, %s", sql, args)
	rows, err := r.conPool.Query(ctx, sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *repository) Delete(ctx context.Context, user *model.Query) error {
	builder := sq.Delete(tableName).PlaceholderFormat(sq.Dollar).
		Where("username = ?", user.Username)

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
