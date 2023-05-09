package server

import (
	"context"

	authpb "github.com/the-spine/spine-protos-go/auth"
)

type server struct{}

func (s *server) Login(context.Context, *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	return &authpb.LoginResponse{}, nil
}

func (s *server) Logout(context.Context, *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {

	return &authpb.LogoutResponse{}, nil
}

func (s *server) RefreshToken(context.Context, *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {

	return &authpb.RefreshTokenResponse{}, nil
}

func (s *server) GetUser(context.Context, *authpb.UserRequest) (*authpb.UserResponse, error) {

	return &authpb.UserResponse{}, nil
}

func (s *server) RegisterUser(context.Context, *authpb.UserRegisterRequest) (*authpb.UserRegisterResponse, error) {

	return &authpb.UserRegisterResponse{}, nil
}

func (s *server) ResetPassword(context.Context, *authpb.PasswordResetRequest) (*authpb.PasswordResetResponse, error) {

	return &authpb.PasswordResetResponse{}, nil
}
