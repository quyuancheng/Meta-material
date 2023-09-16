package service

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	jwtAuth "github.com/golang-jwt/jwt/v5"
	"time"
	pb "verifyCode/api/jwt"
)

type JwtService struct {
	pb.UnimplementedJwtServer
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

var jwtSecretKey *ecdsa.PrivateKey

func (s *JwtService) CreateJwt(ctx context.Context, req *pb.CreateJwtRequest) (*pb.CreateJwtReply, error) {
	curve := elliptic.P256() // choose the curve
	privateKey, _ := ecdsa.GenerateKey(curve, rand.Reader)
	t := jwtAuth.NewWithClaims(jwtAuth.SigningMethodES256,
		jwtAuth.MapClaims{
			"exp":        time.Now().Add(10 * time.Minute),
			"authorized": true,
			"user":       req.UserName,
		})
	tokenString, err := t.SignedString(privateKey)
	if err != nil {
		return &pb.CreateJwtReply{
			Code:  400,
			Token: err.Error(),
		}, nil
	}
	return &pb.CreateJwtReply{
		Code:  200,
		Token: tokenString,
	}, nil
}

func (s *JwtService) GetJwt(ctx context.Context, req *pb.GetJwtRequest) (*pb.GetJwtReply, error) {
	return &pb.GetJwtReply{}, nil
}
