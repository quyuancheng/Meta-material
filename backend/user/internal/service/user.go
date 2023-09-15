package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"regexp"
	"time"
	pb "user/api/user"
	"user/api/verifyCode"
	"user/internal/biz"
)

type UserService struct {
	pb.UnimplementedUserServer
	userUsecase *biz.UserUsecase
}

func NewUserService(userUsecase *biz.UserUsecase) *UserService {
	return &UserService{
		userUsecase: userUsecase,
	}
}

func (us *UserService) GetVerifyCode(ctx context.Context, req *pb.GetVerifyCodeReq) (*pb.GetVerifyCodeResp, error) {
	if !VerifyPhoneFormat(req.Phone) {
		return &pb.GetVerifyCodeResp{
			Code:    400,
			Message: "Phone format error",
		}, nil
	}
	// 通过验证码服务生成验证码 grpc
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9000"),
	)
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    400,
			Message: "The verification code service is not available",
		}, nil
	}
	defer func() {
		_ = conn.Close()
	}()
	// 发送验证码请求
	client := verifyCode.NewVerifyCodeClient(conn)
	response, err := client.GetVerifyCode(context.Background(), &verifyCode.GetVerifyCodeRequest{
		Length: 6,
		Type:   1,
	})
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    400,
			Message: "get verification code failed",
		}, nil
	}
	const life = 60
	err = us.userUsecase.SetVerifyCode(context.Background(), req.Phone, response.Code, life)
	if err != nil {
		return &pb.GetVerifyCodeResp{
			Code:    400,
			Message: "set verification code failed",
		}, nil
	}
	return &pb.GetVerifyCodeResp{
		Code:           200,
		Message:        "success",
		VerifyCode:     response.Code,
		VerifyCodeTime: time.Now().Unix(),
		VerifyCodeLife: life,
	}, nil
}

func VerifyPhoneFormat(phone string) bool {
	// 手机格式校验
	pattern := `^1[3|4|5|7|8][0-9]{9}$`
	regexpPattern := regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(phone) {
		return false
	}
	return true
}
