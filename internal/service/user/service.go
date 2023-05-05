package user

import (
	"context"

	model "github.com/illabb13/user-auth-api/internal/model/user"
	repo "github.com/illabb13/user-auth-api/internal/repository/user"
)

type Service interface {
	Create(ctx context.Context, user *model.UserNewData) error
	Get(ctx context.Context, query *model.Query) (*model.UserInfo, error)
	Delete(ctx context.Context, query *model.Query) error
	Update(ctx context.Context, query *model.Query, user *model.UserNewData) error
}

type service struct {
	repository repo.Repository
}

func NewService(repository repo.Repository) *service {
	return &service{
		repository: repository,
	}
}
