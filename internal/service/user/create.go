package user

import (
	"context"

	"github.com/illabb13/users_auth/internal/model"
)

func (s *service) Create(ctx context.Context, model *model.User) error {
	err := s.repository.Create(ctx, model)
	if err != nil {
		return err
	}
	return nil
}
