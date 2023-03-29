package cache

import (
	"context"
	"encoding/json"
	"forum/internal/model"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) *RedisCache {
	return &RedisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *RedisCache) Set(ctx context.Context, key string, value *model.Post) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		return
	}

	client.Set(ctx, key, json, cache.expires*time.Second)
}

func (cache *RedisCache) Get(ctx context.Context, key string) *model.Post {
	client := cache.getClient()

	value, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	post := model.Post{}

	if err := json.Unmarshal([]byte(value), &post); err != nil {
		return nil
	}

	return &post
}
