package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	models "user/model"
)

func init() {
	models.RegistModel(&User{})
}

// User is a User model.
type User struct {
	ID        uint      `json:"-" gorm:"primary_key;comment:'主键 ID'"`
	CreatedAt time.Time `json:"createTime" gorm:"column:createTime;comment:'创建时间'"`
	UpdatedAt time.Time `json:"updateTime" gorm:"column:updateTime;comment:'更新时间'"`
	Name      string    `json:"name" gorm:"column:name;comment:'用户名称'"`
	Email     string    `json:"email" gorm:"column:email;comment:'邮箱'"`
	Phone     string    `json:"phone" gorm:"column:phone;comment:'手机号'"`
}

// UserRepo is a User repo.
type UserRepo interface {
	SetVerifyCode(ctx context.Context, phone, code string, ex int64) error
	GetVerifyCode(ctx context.Context, phone string) (string, error)
	GetUserByAccount(ctx context.Context, account string) (*User, error)
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

// GetVerifyCode set verify code to redis, and returns status
func (us *UserUsecase) GetVerifyCode(ctx context.Context, phone string) (string, error) {
	us.log.WithContext(ctx).Infof("get verify code")
	verifyCode, err := us.repo.GetVerifyCode(ctx, phone)
	return verifyCode, err
}

// GetUserByAccount -
func (us *UserUsecase) GetUserByAccount(ctx context.Context, account string) (*User, error) {
	return us.repo.GetUserByAccount(ctx, account)
}
