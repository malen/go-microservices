package product

import (
	"aoisoft/gateway/config"
	"aoisoft/gateway/product/pb"
	"fmt"

	"google.golang.org/grpc"
)

type ServiceClient struct {
	pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductService, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect", err)
	}

	return pb.NewProductServiceClient(cc)
}
