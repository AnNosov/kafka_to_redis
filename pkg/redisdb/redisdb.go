package redisdb

import (
	"ktd/config"
	"net"

	"github.com/go-redis/redis/v9"
)

type RedisClient struct {
	RDB *redis.Client
}

func newRedisCfg(cfg config.Redis) *redis.Options {
	return &redis.Options{
		Addr:     net.JoinHostPort(cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0, // default
	}
}

func NewClient(cfg config.Redis) *RedisClient {
	rClient := redis.NewClient(newRedisCfg(cfg))
	rClnt := &RedisClient{
		RDB: rClient,
	}
	return rClnt
}

func (rClient *RedisClient) Close() error {
	return rClient.RDB.Close()
}
