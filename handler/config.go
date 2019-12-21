package handler

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"os"
)
var Conf map[string]interface{}
// InitConfig reads in config file and ENV variables if set.
func InitConfig(cfgFile string)  {
	if cfgFile != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".acm-runner")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	Conf = make(map[string]interface{})
	Conf = viper.AllSettings()
}
