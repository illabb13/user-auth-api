package user

import (
	"context"

	"github.com/illabb13/users_auth/internal/model"
	repo "github.com/illabb13/users_auth/internal/repository/user"
)

type Service interface {
	Create(ctx context.Context, model *model.User) error
	Delete(ctx context.Context, model *model.Query) error
	Update(ctx context.Context, query *model.Query, user *model.User) error
}

type service struct {
	repository repo.Repository
}

func NewService(repository repo.Repository) *service {
	return &service{
		repository: repository,
	}
}
