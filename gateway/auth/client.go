package auth

import (
	"aoisoft/gateway/auth/pb"
	"aoisoft/gateway/config"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthService, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connet:", err)
	}

	return pb.NewAuthServiceClient(cc)
}
