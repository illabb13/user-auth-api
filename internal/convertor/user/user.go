package user

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/illabb13/users_auth/internal/model"
	pkg "github.com/illabb13/users_auth/pkg/user_v1"
)

func ToQuery(search *pkg.QueryInfo) *model.Query {
	return &model.Query{
		Username: search.GetUsername(),
	}
}

func ToUserNewData(user *pkg.UserNewData) *model.UserNewData {
	return &model.UserNewData{
		Username: user.GetUsername(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
		Role:     user.GetRole().String(),
	}
}

func ToUserInfo(user *pkg.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Username:  user.GetUsername(),
		Email:     user.GetEmail(),
		Role:      user.GetRole().String(),
		CreatedAt: user.GetCreatedAt().AsTime(),
		UpdatedAt: user.GetUpdatedAt().AsTime(),
	}
}

func FromUserInfo(userInfo *model.UserInfo) *pkg.UserInfo {
	return &pkg.UserInfo{
		Username:  userInfo.Username,
		Email:     userInfo.Email,
		Role:      pkg.Role(pkg.Role_value[userInfo.Role]),
		CreatedAt: timestamppb.New(userInfo.CreatedAt),
		UpdatedAt: timestamppb.New(userInfo.UpdatedAt),
	}
}
