package user_v1

import (
	service "github.com/illabb13/user-auth-api/internal/service/user"
	pkg "github.com/illabb13/user-auth-api/pkg/user_v1"
)

type Implementation struct {
	pkg.UnimplementedUserV1Server

	service service.Service
}

func NewImplementation(service service.Service) *Implementation {
	return &Implementation{
		service: service,
	}
}
