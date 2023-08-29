package gwebz

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestNewRedisEngine(t *testing.T) {
	redis, err := NewRedisClient("127.0.0.1:6379", "", "123456", 1)
	if err != nil {
		panic(err)
	}
	rsp := "i love you"

	redis.redisClient.Set(context.Background(), "dongzhi", "i love you", 10*time.Second)
	record, err := redis.redisClient.Get(context.Background(), "dongzhi").Result()
	if err != nil {
		panic(err)
	}
	if record != rsp {
		panic(err)
	}
	time.Sleep(10 * time.Second)
	_, err = redis.redisClient.Get(context.Background(), "dongzhi").Result()
	if err == nil {
		panic(errors.New("data not expired"))
	}
}

func TestInitRedisEngine(t *testing.T) {
	err := InitGlobalRedisClient("127.0.0.1:6379", "", "123456", 1)
	if err != nil {
		panic(err)
	}
	rsp := "i love you"
	redisClient := GetRedisClient()
	redisClient.Set(context.Background(), "dongzhi", "i love you", 5*time.Second)
	record, err := redisClient.Get(context.Background(), "dongzhi").Result()
	if err != nil {
		panic(err)
	}
	if record != rsp {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	_, err = redisClient.Get(context.Background(), "dongzhi").Result()
	if err == nil {
		panic(errors.New("data not expired"))
	}
}
