package order

import (
	"aoisoft/gateway/config"
	"aoisoft/gateway/order/pb"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderService, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewOrderServiceClient(cc)
}
