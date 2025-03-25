package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type PublicRedisManager struct {
	Cli        *redis.Client
	ClusterCli *redis.ClusterClient
}

type PublicRedisConfig struct {
	PublicRedisCli
	PublicRedisClu
}

type PublicRedisCli struct {
	Addr     string
	Password string
	DB       int
}

type PublicRedisClu struct {
	Addr     []string
	Password string
}

type PublicRedisFactory struct{}

func NewRedisClient() *PublicRedisFactory {
	return &PublicRedisFactory{}
}

func (f *PublicRedisFactory) CreateRedisManager(config PublicRedisConfig) (*PublicRedisManager, string) {
	var cli *redis.Client
	var clusterCli *redis.ClusterClient

	if len(config.PublicRedisClu.Addr) != 0 {
		clusterCli = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    config.PublicRedisClu.Addr,
			Password: config.PublicRedisClu.Password,
		})
		err := clusterCli.Ping(context.Background()).Err()
		if err != nil {
			return nil, "redis集群连接失败"
		}
	} else {
		cli = redis.NewClient(&redis.Options{
			Addr:     config.PublicRedisCli.Addr,
			Password: config.PublicRedisCli.Password,
			DB:       config.PublicRedisCli.DB,
		})
		err := cli.Ping(context.Background()).Err()
		if err != nil {
			return nil, "redis单机连接失败"
		}
	}

	return &PublicRedisManager{
		Cli:        cli,
		ClusterCli: clusterCli,
	}, ""
}

// Close 关闭 Redis 连接
func (f *PublicRedisFactory) Close(manager PublicRedisManager) error {
	if manager.Cli != nil {
		if err := manager.Cli.Close(); err != nil {
			return err
		}
	}
	if manager.ClusterCli != nil {
		if err := manager.ClusterCli.Close(); err != nil {
			return err
		}
	}
	return nil
}
