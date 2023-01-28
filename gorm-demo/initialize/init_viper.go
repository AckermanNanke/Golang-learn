package initialize

import (
	"flag"
	"fmt"
	"gorm-demo/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper 初始化
// 优先级: 命令行 > 环境变量 > 默认值
func InitViper() *viper.Viper {
	var config string

	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()

	if config == "" {
		if configEnv := os.Getenv(global.ConfigEnv); configEnv == "" { // 判断 全局变量里的 ConfigEnv 常量存储的环境变量是否为空
			config = global.ConfigDefaultFile
			// fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", gin.EnvGinMode, ConfigDefaultFile)
			fmt.Printf("您正在使用gin模式的%s环境名称,config的路径为%s\n", global.ConfigEnv, global.ConfigDefaultFile)
		} else {
			config = configEnv
			fmt.Printf("您正在使用%s环境变量,config的路径为%s\n", global.ConfigEnv, config)
		}
	} else {
		fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(config)

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	v.WatchConfig()

	// 文件发生变化（仅检测是否覆盖了不检测内容）时调用
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		fmt.Println(err)
	}

	return v
}
