package user_v1

import (
	"context"
	"log"

	converter "github.com/illabb13/users_auth/internal/convertor/user"
	pkg "github.com/illabb13/users_auth/pkg/user_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *pkg.DeleteRequest) (*pkg.DeleteResponse, error) {
	query := converter.ToQuery(req.GetQuery())
	log.Printf("deleting user by params: %s\n", query)
	err := i.service.Delete(ctx, query)
	if err != nil {
		return nil, err
	}
	return &pkg.DeleteResponse{}, nil
}
