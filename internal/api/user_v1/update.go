package user_v1

import (
	"context"
	"log"

	converter "github.com/illabb13/user-auth-api/internal/convertor/user"
	pkg "github.com/illabb13/user-auth-api/pkg/user_v1"
)

func (i *Implementation) Update(ctx context.Context, req *pkg.UpdateRequest) (*pkg.UpdateResponse, error) {
	query := converter.ToQuery(req.GetQuery())
	user := converter.ToUserNewData(req.GetUser())
	log.Printf("updating user by query: %s, %s\n", user, query)
	err := i.service.Update(ctx, query, user)
	if err != nil {
		return nil, err
	}
	return &pkg.UpdateResponse{}, nil
}
