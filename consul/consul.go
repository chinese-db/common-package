package consul

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

// Client 封装了Consul客户端的操作
type Client struct {
	client *api.Client
}

// NewClient 创建新的Consul客户端实例
// host: Consul服务器地址（包含端口，如 "localhost:8500"）
func NewClient(host string) (*Client, error) {
	config := api.DefaultConfig()
	config.Address = host

	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Consul client: %w", err)
	}

	return &Client{client: client}, nil
}

// ServiceRegistration 服务注册参数
type ServiceRegistration struct {
	Name    string
	Address string
	Port    int
	Tags    []string
	ID      string // 可选，如果为空则自动生成
}

// RegisterService 注册服务到Consul
func (c *Client) RegisterService(reg ServiceRegistration) (string, error) {
	if reg.ID == "" {
		reg.ID = uuid.NewString()
	}

	registration := &api.AgentServiceRegistration{
		ID:      reg.ID,
		Name:    reg.Name,
		Address: reg.Address,
		Port:    reg.Port,
		Tags:    reg.Tags,
	}

	err := c.client.Agent().ServiceRegister(registration)
	if err != nil {
		return "", fmt.Errorf("failed to register service: %w", err)
	}

	return reg.ID, nil
}

// DeregisterService 注销服务
func (c *Client) DeregisterService(serviceID string) error {
	err := c.client.Agent().ServiceDeregister(serviceID)
	if err != nil {
		return fmt.Errorf("failed to deregister service: %w", err)
	}
	return nil
}

// DiscoverServices 服务发现
// filter: 过滤表达式，如 "Service == web"
func (c *Client) DiscoverServices(filter string) (map[string]*api.AgentService, error) {
	services, err := c.client.Agent().ServicesWithFilter(filter)
	if err != nil {
		return nil, fmt.Errorf("failed to discover services: %w", err)
	}
	return services, nil
}

// GetAllServices 获取所有注册的服务
func (c *Client) GetAllServices() (map[string]*api.AgentService, error) {
	services, err := c.client.Agent().Services()
	if err != nil {
		return nil, fmt.Errorf("failed to get all services: %w", err)
	}
	return services, nil
}

// HealthCheck 添加健康检查（示例方法）
func (c *Client) AddHealthCheck(check *api.AgentCheckRegistration) error {
	return c.client.Agent().CheckRegister(check)
}
