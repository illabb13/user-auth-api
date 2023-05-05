package user

import (
	"context"

	model "github.com/illabb13/users_auth/internal/model/user"
)

func (s *service) Create(ctx context.Context, user *model.UserNewData) error {
	err := s.repository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
