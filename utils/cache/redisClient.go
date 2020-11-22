package cache

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
)

type Cache struct {
	client *redis.Client
}

func (c *Cache) Save(key string, value string, expired time.Duration) error {
	err := c.client.Set(key, value, expired).Err()

	fmt.Println(err)

	return err
}

func (c *Cache) Read(key string) (string, error) {
	return c.client.Get(key).Result()
}

func (c *Cache) Delete(key string) error {
	return c.client.Del(key).Err()
}

func redisInit() *Cache {
	cache := &Cache{
		client: redis.NewClient(&redis.Options{
			Addr:     os.Getenv("CACHE_CONNECTION_STRING"),
			Password: os.Getenv("CACHE_PASSWORD"),
			DB:       0,
		}),
	}

	fmt.Println("we says : ping")
	pong, err := cache.client.Ping().Result()

	fmt.Println("redis says : " + pong)

	if err != nil {
		panic(err)
	}

	return cache
}
