package usecase

import (
	"context"
	"fmt"
	"ktd/internal/entity"
	"ktd/pkg/redisdb"
	"log"
	"strconv"

	"github.com/go-redis/redis/v9"
)

type RedisProfile struct {
	*redisdb.RedisClient
}

func NewRedisClient(r *redisdb.RedisClient) *RedisProfile {
	return &RedisProfile{r}
}

func (r *RedisProfile) PutProfiles(profileChan chan entity.Profile) {
	profile := <-profileChan

	if err := r.putProfileHash(profile); err != nil {
		log.Println("Redis PutProfiles ", err)
		return
	} else if err := r.putProfileKeys(strconv.Itoa(profile.Id)); err != nil {
		log.Println("Redis PutProfiles ", err)
		err := r.deleteKeyProfile(strconv.Itoa(profile.Id))
		if err != nil {
			log.Println("Redis PutProfiles ", err)
		}
		return
	}
}

func (r *RedisProfile) putProfileKeys(key string) error {
	//  ключ ключей - prfls
	err := r.RedisClient.RDB.Set(context.Background(), "prfls", key, redis.KeepTTL).Err()
	if err != nil {
		return fmt.Errorf("putProfileKeys %w", err)
	}
	return nil
}

func (r *RedisProfile) putProfileHash(prfl entity.Profile) error {
	err := r.RedisClient.RDB.HSet(context.Background(), strconv.Itoa(prfl.Id), prfl).Err()
	if err != nil {
		return fmt.Errorf("putProfileHash %w", err)
	}
	return nil
}

func (r *RedisProfile) deleteKeyProfile(key string) error {
	if err := r.RedisClient.RDB.SRem(context.Background(), "prfls", key).Err(); err != nil {
		return fmt.Errorf("deleteKeyProfile %w", err)
	}
	return nil
}
