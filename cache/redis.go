package cache

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
	"github.com/todzuko/go-api/models"
	"time"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) QuestsCache {
	return &redisCache{
		host,
		db,
		exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *models.Quest) {
	client := cache.getClient()
	jsonVal, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key, jsonVal, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *models.Quest {
	client := cache.getClient()
	jsonVal, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	quest := models.Quest{}
	err = json.Unmarshal([]byte(jsonVal), &quest)
	if err != nil {
		panic(err)
	}
	return &quest
}
