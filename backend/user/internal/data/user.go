package data

// user中与数据操作相关的代码
import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
	"user/internal/biz"
)

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{data: data, log: log.NewHelper(logger)}
}

func (u *UserRepo) SetVerifyCode(ctx context.Context, phone, code string, ex int64) error {
	// 设置key
	status := u.data.Rdb.Set(ctx, phone, code, time.Duration(ex)*time.Second)
	if _, err := status.Result(); err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) GetVerifyCode(ctx context.Context, phone string) (string, error) {
	// 获取验证码
	res := u.data.Rdb.Get(ctx, phone)
	if code, err := res.Result(); err != nil {
		return "", err
	} else {
		return code, nil
	}
}

func (u *UserRepo) GetUserByAccount(ctx context.Context, account string) (*biz.User, error) {
	var user *biz.User
	err := u.data.DB.Where("phone", account).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
