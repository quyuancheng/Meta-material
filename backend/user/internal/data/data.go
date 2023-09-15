package data

import (
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"user/internal/conf"
	models "user/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo, NewUserRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Rdb *redis.Client
	DB  *gorm.DB
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
	// 初始化mysql
	dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", c.Database.Source)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		data.DB = nil
	}
	// 建立mysql客户端链接
	data.DB = db
	// 迁移表
	tables := models.GetModels()
	tableOptions := "ENGINE=InnoDB DEFAULT CHARSET=utf8"
	for _, table := range tables {
		err = db.Set("gorm:table_options", tableOptions).AutoMigrate(table)
		if err != nil {
			fmt.Errorf("auto migrate table %+v failure %s", table, err.Error())
			return nil, nil, err
		}
	}
	cleanup := func() {
		_ = data.Rdb.Close()
		log.NewHelper(logger).Info("closing the data resources")
	}
	return data, cleanup, nil
}
