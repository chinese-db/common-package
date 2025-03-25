package nacos

import (
	"encoding/json"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// InitNacos 初始化 Nacos 客户端并加载配置
func InitNacos(configPath string) error {
	// 加载配置文件
	if err := LoadConfig(configPath); err != nil {
		return err
	}

	// 创建 ClientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         Conf.Nacos.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建 ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      Conf.Nacos.Addr,
			ContextPath: "/nacos",
			Port:        Conf.Nacos.Port,
			Scheme:      "http",
		},
	}

	// 创建 Nacos 客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return err
	}

	// 获取配置内容
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: Conf.Nacos.Dataid,
		Group:  Conf.Nacos.Group,
	})
	if err != nil {
		return err
	}

	// 解析配置内容
	var appConf map[string]interface{}
	if err := json.Unmarshal([]byte(content), &appConf); err != nil {
		return err
	}

	// 将解析后的配置存储到全局变量中
	//Conf.AppConf = appConf
	return nil
}
