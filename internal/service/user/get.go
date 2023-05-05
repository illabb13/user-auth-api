package user

import (
	"context"

	"github.com/illabb13/users_auth/internal/model"
)

func (s *service) Get(ctx context.Context, query *model.Query) (*model.UserInfo, error) {
	user, err := s.repository.Get(ctx, query)
	if err != nil {
		return nil, err
	}
	return user, nil
}
