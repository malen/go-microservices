package services

import (
	"aoisoft/gateway/order/pb"
	"aoisoft/order/client"
	"aoisoft/order/models"
	"aoisoft/order/models/db"
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	H              db.Handler
	ProductService client.ProductServiceClient
	pb.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	fmt.Println("CreateOrder start")
	product, err := s.ProductService.FindOne(req.ProductId)
	fmt.Println(product)
	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if product.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{Status: product.Status, Error: product.Error}, nil
	} else if product.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: "Stock too less"}, nil
	}

	order := models.Order{
		Price:     product.Data.Price,
		ProductId: product.Data.Id,
		UserId:    req.UserId,
	}

	s.H.DB.Insert(&order)

	res, err := s.ProductService.DecreaseStock(req.ProductId, order.Id)
	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.Id)
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
