package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

// InitNacos 初始化 Nacos 客户端并加载配置
func InitNacos(configPath string) (string, error) {
	// 调用LoadConfig加载配置文件
	nacos, err := LoadConfig(configPath)
	if err != nil {
		panic("nacos配置解析失败")
	}
	// 创建 ClientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacos.Nacos.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建 ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacos.Nacos.Addr,
			ContextPath: "/nacos",
			Port:        nacos.Nacos.Port,
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
		return "", err
	}

	// 获取配置内容
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: nacos.Nacos.Dataid,
		Group:  nacos.Nacos.Group,
	})
	if err != nil {
		return "", err
	}

	return content, nil
}
