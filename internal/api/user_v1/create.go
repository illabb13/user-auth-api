package user_v1

import (
	"context"
	"log"

	converter "github.com/illabb13/users_auth/internal/convertor/user"
	pkg "github.com/illabb13/users_auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *pkg.CreateRequest) (*pkg.CreateResponse, error) {
	user := converter.ToUser(req.GetUser())
	log.Printf("creating user: %s\n", user)
	err := i.service.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &pkg.CreateResponse{}, nil
}
