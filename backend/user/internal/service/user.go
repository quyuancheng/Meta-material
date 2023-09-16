package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"regexp"
	"time"
	"user/api/jwt"
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

func (us *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	//TODO: 先只验证手机号，后续会增加邮箱和其他的
	if !VerifyPhoneFormat(req.Account) {
		return &pb.LoginResp{
			Code:  400,
			Token: "",
		}, nil
	}
	// 查询手机号是否存在，不存在提示注册
	user, err := us.userUsecase.GetUserByAccount(context.Background(), req.Account)
	if err != nil {
		return &pb.LoginResp{
			Code:    400,
			Token:   "",
			Message: "Phone number does not exist, please register",
		}, nil
	}
	// 查询密码是否正确
	// 查询验证码是否过期，如果过期则提示验证码过期
	code, err := us.userUsecase.GetVerifyCode(context.Background(), req.Account)
	if code == "" && err != nil {
		return &pb.LoginResp{
			Code:    400,
			Token:   "",
			Message: "The verification code has expired",
		}, nil
	}
	// jwt服务
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("localhost:9000"),
	)
	if err != nil {
		return &pb.LoginResp{
			Code:    400,
			Message: "Failed to generate token",
		}, nil
	}
	defer func() {
		_ = conn.Close()
	}()
	// 生成jwt
	client := jwt.NewJwtClient(conn)
	j, err := client.CreateJwt(context.Background(), &jwt.CreateJwtRequest{
		UserName: user.Name,
	})
	if err != nil {
		return &pb.LoginResp{
			Code:    400,
			Message: "create jwt failed",
		}, nil
	}
	// 登录成功，返回token
	return &pb.LoginResp{
		Code:    200,
		Token:   j.Token,
		Message: "login success",
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
