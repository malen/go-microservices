package services

import (
	"aoisoft/gateway/product/pb"
	"aoisoft/product/models"
	"aoisoft/product/models/db"
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	H db.Handler
	pb.UnimplementedProductServiceServer
}

func (s *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	var product = models.Product{
		Name:  req.Name,
		Stock: req.Stock,
		Price: req.Price,
	}

	if _, err := s.H.DB.Insert(&product); err != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var products []models.Product

	if err := s.H.DB.Where("id = ?", req.Id).Find(&products); err != nil {
		fmt.Println("xxxxxxxxxxxxx", err)
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    products[0].Id,
		Name:  products[0].Name,
		Stock: products[0].Stock,
		Price: products[0].Price,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

// FIXME: 需要开启事务
func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var product models.Product
	if _, err := s.H.DB.ID(req.Id).Get(&product); err != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, nil
	}

	if product.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var logs []models.StockDecreaseLog
	if err := s.H.DB.Where("orderid = ?", req.OrderId).Find(&logs); err != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	product.Stock = product.Stock - 1
	s.H.DB.Update(&product)

	var log = models.StockDecreaseLog{
		OrderId:      req.OrderId,
		ProductRefer: product.Id,
	}
	s.H.DB.Insert(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
