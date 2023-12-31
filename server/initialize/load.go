package initialize

import (
	"vue-admin-element/global"
	"fmt"

	"github.com/spf13/viper"
)

// 加载配置文件
func LoadConfig() {
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Fatal error resources file: %s \n", err.Error())
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("unable to decode into struct %s \n", err.Error())
	}
}
