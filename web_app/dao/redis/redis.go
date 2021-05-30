package redis

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/go-redis/redis"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化redis连接
func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),   //连接池大小
	})
	_, err = rdb.Ping().Result()
	return
}

// Close 关闭redis连接
func Close() {
	_ = rdb.Close()
}
