package server

import (
	"auth/internal/models"
	"auth/internal/repository"
	"context"

	authpb "github.com/the-spine/spine-protos-go/auth"
)

type authServer struct {
	authpb.UnimplementedAuthServiceServer
}

func GetAuthServer() *authServer {
	return &authServer{}
}

func (s *authServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	return &authpb.LoginResponse{}, nil
}

func (s *authServer) Logout(ctx context.Context, req *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {

	return &authpb.LogoutResponse{}, nil
}

func (s *authServer) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {

	return &authpb.RefreshTokenResponse{}, nil
}

func (s *authServer) GetUser(ctx context.Context, req *authpb.UserRequest) (*authpb.UserResponse, error) {

	return &authpb.UserResponse{}, nil
}

func (s *authServer) RegisterUser(ctx context.Context, req *authpb.UserRegisterRequest) (*authpb.UserRegisterResponse, error) {

	user := models.User{
		FirstName:    req.GetFirstName(),
		MiddleName:   req.GetMiddleName(),
		LastName:     req.GetLastName(),
		Email:        req.GetEmail(),
		PasswordHash: req.GetPassword(),
	}

	err := repository.CreateUser(&user)

	if err != nil {
		return &authpb.UserRegisterResponse{Success: false}, err
	}

	return &authpb.UserRegisterResponse{Success: true}, nil
}

func (s *authServer) ResetPassword(ctx context.Context, req *authpb.PasswordResetRequest) (*authpb.PasswordResetResponse, error) {

	return &authpb.PasswordResetResponse{}, nil
}
