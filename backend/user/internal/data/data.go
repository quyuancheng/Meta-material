package data

import (
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	data := &Data{}
	// 初始化redis
	redisURL := fmt.Sprintf("redis://%s/1?protocol=3&dial_timeout=1", c.Redis.Addr)
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		data.Rdb = nil
	}
	// 建立redis客户端链接
	data.Rdb = redis.NewClient(options)
	cleanup := func() {
		_ = data.Rdb.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
