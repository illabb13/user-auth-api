package user

import (
	"github.com/illabb13/users_auth/internal/model"
	pkg "github.com/illabb13/users_auth/pkg/user_v1"
)

func ToQuery(search *pkg.QueryInfo) *model.Query {
	return &model.Query{
		Username: search.GetUsername(),
	}
}

func ToUser(user *pkg.UserInfo) *model.User {
	return &model.User{
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Role:     user.GetRole(),
	}
}
