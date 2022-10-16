package services

import (
	"aoisoft/auth/models"
	"aoisoft/auth/models/db"
	"aoisoft/auth/utils"
	"aoisoft/gateway/auth/pb"
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var users []models.User
	if err := s.H.DB.Where("email = ?", req.Email).Find(&users); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}

	var user models.User
	user.Email = req.Email
	user.Password = utils.HashPassword(req.Password)
	s.H.DB.Insert(&user)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var users []models.User
	if err := s.H.DB.Where("email = ?", req.Email).Find(&users); err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "user not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, users[0].Password)
	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(users[0])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("xxxx %s\n", token)
	fmt.Println(users[0])
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.Jwt.ValidateToken(req.Token)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	var users []models.User
	if err := s.H.DB.Where("email = ?", claims.Email).Find(&users); err != nil {
		fmt.Printf("xxxxxxxxx [%s]\n", claims.Email)
		fmt.Println(err)
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: users[0].Id,
	}, nil
}
