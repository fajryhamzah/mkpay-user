package cache

import "sync"

var once sync.Once
var instance *Cache

func Get() Cache {
	once.Do(func() {
		instance = redisInit()
	})

	return *instance
}
