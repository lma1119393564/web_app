package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Conf 将所有配置信息存放到和结构体指针中
var Conf = new(AppConfig)

type AppConfig struct {
	Name  string `mapstructure:"name"`//
	Mode	string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port	int `mapstructure:"port"`
	*LogConfig `mapstructure:"log"` //log:
	*MysqlConfig `mapstructure:"mysql"` //mysql:
	*RedisConfig `mapstructure:"redis"` //redis:
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int `mapstructure:"max_size"`
	MaxAge     int `mapstructure:"max_age"`
	MaxBackups int `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Dbname       string `mapstructure:"dbname"`
	Port         int `mapstructure:"port"`
	MaxOpenConns int `mapstructure:"max_open_conns"`
	MaxIdleConns int `mapstructure:"max_idle_conns"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}




func Init() (err error) {
	//viper.SetConfigFile("E:/gowork/src/qimi_web/web_app2/config.yaml") // 指定配置文件路径(绝对路径)
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径（相对路径）

	//viper.SetConfigName("config") // 配置文件名称(无扩展名)
	//viper.SetConfigType("yaml")   // 从远程源返回配置类型
	//viper.AddConfigPath(".")      // 查找配置文件所在的路径
	//viper.AddConfigPath("$HOME/.appname") // 多次调用以添加多个搜索路径
	//viper.AddConfigPath(".")              // 还可以在工作目录中查找配置
	err = viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {            // 处理读取配置文件的错误
		fmt.Printf("viper.ReadInConfig() failed,err:=%v\n", err)
		return
	}
	//配置信息反序列化到 conf 结构体中
	err = viper.Unmarshal(Conf)
	if err != nil {
		fmt.Printf("viper.Unmarshal(Conf) failed,err:%v\n",err)
	}
	//监视配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
		// 改变后在进行序列化
		if err = viper.Unmarshal(Conf); err!=nil{
			fmt.Printf("viper.Unmarshal(Conf) failed,err:%v\n",err)
		}
	})
	return
}
