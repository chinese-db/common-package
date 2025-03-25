package nacos

import (
	"github.com/spf13/viper"
)

// 配置结构体
type Config struct {
	Nacos struct {
		Namespace string `mapstructure:"namespace"`
		Addr      string `mapstructure:"addr"`
		Port      uint64 `mapstructure:"port"`
		Dataid    string `mapstructure:"dataid"`
		Group     string `mapstructure:"group"`
	} `mapstructure:"nacos"`
}

var Conf Config

// LoadConfig 加载配置文件
func LoadConfig(configPath string) error {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return viper.Unmarshal(&Conf)
}
