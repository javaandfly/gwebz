package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var _redisClient *redis.Client

func GetRedisClient() *redis.Client {
	return _redisClient
}

func SetRedisClient(redisClient *redis.Client) {
	_redisClient = redisClient
}

type KVEngine struct {
	redisClient *redis.Client
}

func InitKVEngine(addr, username, password string, db int) (*KVEngine, error) {
	opt := &redis.Options{
		Addr:     addr,     // host:port address.
		Username: username, // username
		Password: password, // password
		DB:       db,       // Database to be selected after connecting to the server.
	}

	client := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// check connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return &KVEngine{redisClient: client}, err
}

func InitGlobalRedisClient(addr, username, password string, db int) error {
	kvEngine, err := InitKVEngine(addr, username, password, db)
	if err != nil {
		return err
	}
	SetRedisClient(kvEngine.redisClient)
	return nil
}

func NewRedisClient(addr, username, password string, db int) (*KVEngine, error) {
	return InitKVEngine(addr, username, password, db)

}
