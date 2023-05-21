package server

import (
	"auth/internal/config"
	"auth/internal/models"
	"auth/internal/repository"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	authpb "github.com/the-spine/spine-protos-go/auth"
)

type authServer struct {
	config config.Config
	authpb.UnimplementedAuthServiceServer
}

func GetAuthServer(config config.Config) *authServer {
	return &authServer{config: config}
}

func (s *authServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {

	var user models.User

	err := repository.GetUserByEmail(req.GetUsername(), &user)

	if err != nil {
		return nil, err
	}

	exp := time.Now().Add(time.Hour * 24 * 30).UnixMicro()

	// TODO: Implement jti with redis for invocation
	refreshClaims := jwt.MapClaims{
		"iss": "Auth Service",
		"aud": user.TenantID,
		"sub": user.ID,
		"nbf": time.Now().UnixMicro(),
		"exp": exp,
		"jti": "",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims)

	signedRefreshToken, err := refreshToken.SignedString(s.config.JWT.EcdsaPrivateKey)

	if err != nil {
		return nil, err
	}

	acessTokenClaims := jwt.MapClaims{
		"iss":   "Auth Service",
		"aud":   user.TenantID,
		"sub":   user.ID,
		"nbf":   time.Now().UnixMicro(),
		"exp":   time.Now().Add(time.Minute * 30).UnixMicro(),
		"jti":   "",
		"roles": user.Roles,
	}

	acessToken := jwt.NewWithClaims(jwt.SigningMethodES256, acessTokenClaims)

	signedAcessToken, err := acessToken.SignedString(s.config.JWT.EcdsaPrivateKey)

	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		RefreshToken: signedRefreshToken,
		TokenType:    "OAuth 2.0",
		ExpiresIn:    exp,
		AccessToken:  signedAcessToken,
	}, nil
}

func (s *authServer) Logout(ctx context.Context, req *authpb.LogoutRequest) (*authpb.LogoutResponse, error) {

	return &authpb.LogoutResponse{}, nil
}

func (s *authServer) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.RefreshTokenResponse, error) {

	parsedToken, err := jwt.Parse(req.GetRefreshToken(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.config.JWT.EcdsaPublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("couldn't verify vlaims from token")
	}

	now := time.Now().UnixMicro()

	nbf, ok := claims["nbf"].(int64)

	if !ok {
		return nil, fmt.Errorf("couldn't verify vlaims from token")
	}

	if now < nbf {
		return nil, fmt.Errorf("token not yet valid")
	}

	exp, ok := claims["exp"].(int64)

	if !ok {
		return nil, fmt.Errorf("couldn't verify vlaims from token")
	}

	if now > exp {
		return nil, fmt.Errorf("token expired")
	}

	sub, ok := claims["sub"].(uuid.UUID)

	if !ok {
		return nil, fmt.Errorf("couldn't verify claims from token")
	}

	// TODO Make this more efficient and reduce call to DB

	var user models.User

	err = repository.GetUserById(sub.String(), &user)

	if err != nil {
		return nil, fmt.Errorf("token doesn't belong to any user")
	}

	expiration := time.Now().Add(time.Minute * 30).UnixMicro()

	acessTokenClaims := jwt.MapClaims{
		"iss":   "Auth Service",
		"aud":   user.TenantID,
		"sub":   user.ID,
		"nbf":   time.Now().UnixMicro(),
		"exp":   expiration,
		"jti":   "",
		"roles": user.Roles,
	}

	acessToken := jwt.NewWithClaims(jwt.SigningMethodES256, acessTokenClaims)

	signedAcessToken, err := acessToken.SignedString(s.config.JWT.EcdsaPrivateKey)

	if err != nil {
		return nil, err
	}

	return &authpb.RefreshTokenResponse{
		TokenType:   "Acess Token",
		AccessToken: signedAcessToken,
		ExpiresIn:   expiration,
	}, nil
}

func (s *authServer) GetUser(ctx context.Context, req *authpb.UserRequest) (*authpb.UserResponse, error) {

	parsedToken, err := jwt.Parse(req.GetToken(), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.config.JWT.EcdsaPublicKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("couldn't verify vlaims from token")
	}

	now := time.Now().UnixMicro()

	nbf, ok := claims["nbf"].(int64)

	if !ok {
		return nil, fmt.Errorf("couldn't verify vlaims from token")
	}

	if now < nbf {
		return nil, fmt.Errorf("token not yet valid")
	}

	exp, ok := claims["exp"].(int64)

	if !ok {
		return nil, fmt.Errorf("couldn't verify vlaims from token")
	}

	if now > exp {
		return nil, fmt.Errorf("token expired")
	}

	sub, ok := claims["sub"].(uuid.UUID)

	if !ok {
		return nil, fmt.Errorf("couldn't verify claims from token")
	}

	var user models.User

	err = repository.GetUserById(sub.String(), &user)

	if err != nil {
		return nil, fmt.Errorf("token doesn't belong to any user")
	}

	var rolesResponse []*authpb.Role

	for _, v := range user.Roles {
		rolesResponse = append(rolesResponse, &authpb.Role{Role: v.Role, Permissions: v.Permissions})
	}

	// TODO implement userName
	return &authpb.UserResponse{
		UserId:     user.ID.String(),
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		UserName:   user.Email,
		Email:      user.Email,
		Roles:      rolesResponse,
	}, nil
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

	// TODO: generate a email with password reset link

	return &authpb.PasswordResetResponse{}, nil
}
