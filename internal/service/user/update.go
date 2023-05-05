package user

import (
	"context"

	model "github.com/illabb13/user-auth-api/internal/model/user"
)

func (s *service) Update(ctx context.Context, query *model.Query, user *model.UserNewData) error {
	err := s.repository.Update(ctx, query, user)
	if err != nil {
		return err
	}
	return nil
}
