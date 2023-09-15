package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

// User is a User model.
type User struct {
}

// UserRepo is a User repo.
type UserRepo interface {
	SetVerifyCode(ctx context.Context, phone, code string, ex int64) error
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// SetVerifyCode set verify code to redis, and returns status
func (us *UserUsecase) SetVerifyCode(ctx context.Context, phone, code string, ex int64) error {
	us.log.WithContext(ctx).Infof("get verify code")
	status := us.repo.SetVerifyCode(ctx, phone, code, ex)
	return status
}
