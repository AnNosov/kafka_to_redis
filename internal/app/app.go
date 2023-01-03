package app

import (
	"ktd/config"
	"ktd/internal/usecase"
	"ktd/pkg/kfk"
	"ktd/pkg/redisdb"
)

func Run(cfg *config.Config) {
	kcl := kfk.NewReader(cfg.Kafka)
	rcl := redisdb.NewClient(cfg.Redis)
	uc := usecase.New(*usecase.NewKafkaReader(kcl), *usecase.NewRedisClient(rcl))

	uc.TransportData()
}
