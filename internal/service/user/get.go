package user

import (
	"context"

	model "github.com/illabb13/user-auth-api/internal/model/user"
)

func (s *service) Get(ctx context.Context, query *model.Query) (*model.UserInfo, error) {
	user, err := s.repository.Get(ctx, query)
	if err != nil {
		return nil, err
	}
	return user, nil
}
