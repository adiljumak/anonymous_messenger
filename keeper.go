package main

import (
	"context"

	"github.com/go-redis/redis/v8"
)

const NotFoundError = "not_found"

type Keeper interface {
	Get(key string) (string, error)
	Set(key string, message string) error
	Clean(key string) error
}

func getDummyKeeper() Keeper {
	return DummyKeeper{make(map[string]string)}
}

func getRedisKeeper() Keeper {
	return RedisKeeper{
		cn: *redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
		ctx: context.Background(),
	}

}
