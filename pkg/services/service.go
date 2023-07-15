package services

import (
	"context"
	"net/http"

	"github.com/stebinsabu13/note_taking_microservice/auth_srv/pkg/db"
	"github.com/stebinsabu13/note_taking_microservice/auth_srv/pkg/models"
	"github.com/stebinsabu13/note_taking_microservice/auth_srv/pkg/pb"
	"github.com/stebinsabu13/note_taking_microservice/auth_srv/pkg/utils"
)

type Server struct {
	H   db.Handler
	JWT utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User
	if result := s.H.DB.Where("email=?", req.Email).First(&user); result.Error == nil {
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-Mail already exists",
		}, nil
	}
	hash := utils.HashPassword(req.Password)
	user.Name = req.Name
	user.Email = req.Email
	user.Password = hash
	if err := s.H.DB.Create(&user).Error; err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User
	if err := s.H.DB.Where("email=?", req.Email).First(&user).Error; err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "User not exsists",
		}, nil
	}
	match := utils.CheckPasswordHash(req.Password, user.Password)
	if !match {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "invalid password",
		}, nil
	}
	token, err := s.JWT.GenerateToken(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "error generating token",
		}, nil
	}
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.JWT.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: claims.Id,
	}, nil
}
