package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

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
	//监视配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
	})
	return
}
