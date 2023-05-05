package user

import (
	"context"

	"github.com/illabb13/users_auth/internal/model"
)

func (s *service) Update(ctx context.Context, query *model.Query, user *model.UserNewData) error {
	err := s.repository.Update(ctx, query, user)
	if err != nil {
		return err
	}
	return nil
}
