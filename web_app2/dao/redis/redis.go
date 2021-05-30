package redis

import (
	"fmt"
	"qimi_web/web_app/settings"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化redis连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.Db,          // use default DB
		PoolSize: cfg.PoolSize,   //连接池大小
	})
	_, err = rdb.Ping().Result()
	return
}

// Close 关闭redis连接
func Close() {
	_ = rdb.Close()
}
