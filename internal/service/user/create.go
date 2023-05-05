package user

import (
	"context"

	"github.com/illabb13/users_auth/internal/model"
)

func (s *service) Create(ctx context.Context, user *model.User) error {
	err := s.repository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
