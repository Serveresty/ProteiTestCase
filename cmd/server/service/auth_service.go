package service

import (
	"context"
	"fmt"
	"proteitestcase/internal/config"
	"proteitestcase/pkg/api"

	"github.com/dgrijalva/jwt-go"
)

type AuthServer struct {
	api.UnimplementedAuthServiceServer
}

func (s *AuthServer) Login(ctx context.Context, in *api.LoginRequest) (*api.LoginResponce, error) {
	login, password, err := config.GetAuthData()
	if err != nil {
		return nil, err
	}

	if login == in.Login && IsCorrectPassword(password, in.Password) {
		token, err := CreateToken(in.Login)
		if err != nil {
			return &api.LoginResponce{Token: ""}, fmt.Errorf("Error while creating token")
		}
		return &api.LoginResponce{Token: token}, nil
	}

	return &api.LoginResponce{Token: ""}, fmt.Errorf("Bad credentials")
}

func CheckAuth(ctx context.Context) (username string) {
	tokenStr, err := getTokenFromContext(ctx)
	if err != nil {
		panic("get token from context error")
	}
	var clientClaims Claims
	token, err := jwt.ParseWithClaims(tokenStr, &clientClaims, func(token *jwt.Token) (interface{}, error) {
		if token.Header["alg"] != "HS256" {
			panic("ErrInvalidAlgorithm")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		panic("jwt parse error")
	}

	if !token.Valid {
		panic("ErrInvalidToken")
	}

	return clientClaims.Login
}