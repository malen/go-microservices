package main

import (
	"aoisoft/gateway/order/pb"
	"aoisoft/order/client"
	"aoisoft/order/config"
	"aoisoft/order/models/db"
	"aoisoft/order/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)
	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listening", err)
	}

	productService := client.InitProductServiceClient(c.ProductService)

	s := services.Server{
		H:              h,
		ProductService: productService,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve", err)
	}
}
