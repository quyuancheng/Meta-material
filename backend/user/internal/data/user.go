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
