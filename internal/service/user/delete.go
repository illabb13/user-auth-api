package user

import (
	"context"

	"github.com/illabb13/users_auth/internal/model"
)

func (s *service) Delete(ctx context.Context, model *model.Query) error {
	err := s.repository.Delete(ctx, model)
	if err != nil {
		return err
	}
	return nil
}
