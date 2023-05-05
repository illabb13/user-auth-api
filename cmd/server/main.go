package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"

	api "github.com/illabb13/users_auth/internal/api/user_v1"
	repo "github.com/illabb13/users_auth/internal/repository/user"
	service "github.com/illabb13/users_auth/internal/service/user"
	pkg "github.com/illabb13/users_auth/pkg/user_v1"
)

const grpcPort = ":50051"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to get listener: %s", err.Error())
	}

	server := grpc.NewServer()
	reflection.Register(server)

	dbDsn := os.Getenv("PG_DSN")
	dbCfg, err := pgxpool.ParseConfig(dbDsn)
	if err != nil {
		log.Fatalf("failed to get database config: %s", err.Error())
	}

	ctx := context.Background()
	dcCon, err := pgxpool.ConnectConfig(ctx, dbCfg)
	if err != nil {
		log.Fatalf("failed to get database connection: %s", err.Error())
	}
	defer dcCon.Close()

	userRepository := repo.NewRepository(dcCon)
	userService := service.NewService(userRepository)
	pkg.RegisterUserV1Server(server, api.NewImplementation(userService))

	log.Println("start serving")
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
