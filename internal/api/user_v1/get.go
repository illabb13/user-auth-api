package user_v1

import (
	"context"
	"log"

	converter "github.com/illabb13/user-auth-api/internal/convertor/user"
	pkg "github.com/illabb13/user-auth-api/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, req *pkg.GetRequest) (*pkg.GetResponse, error) {
	query := converter.ToQuery(req.GetQuery())
	log.Printf("getting user by query: %s\n", query)
	userInfo, err := i.service.Get(ctx, query)
	if err != nil {
		return nil, err
	}
	result := converter.FromUserInfo(userInfo)
	return &pkg.GetResponse{
		User: result,
	}, nil
}
