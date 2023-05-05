package user

import (
	"context"

	model "github.com/illabb13/users_auth/internal/model/user"
)

func (s *service) Delete(ctx context.Context, user *model.Query) error {
	err := s.repository.Delete(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
