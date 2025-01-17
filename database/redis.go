package database

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type Redis struct {
	client      *redis.Client
	context     context.Context
	Options     *redis.Options
	IsConnected bool
}

func (r *Redis) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

func (r *Redis) Init() error {
	r.context = context.Background()
	r.client = redis.NewClient(r.Options)

	if r.CheckIfConnected() == true {
		return errors.New("unable to connect to Redis")
	}

	return nil
}

// CheckIfConnected Checks our connection status to our Redis DB
func (r *Redis) CheckIfConnected() bool {
	if pong, err := r.client.Ping(r.context).Result(); err != nil && pong == "PONG" {
		r.IsConnected = true
	}
	return r.IsConnected
}

// Get a key value pair from our Redis DB
func (r *Redis) Get(key string) (string, error) {
	value, err := r.client.Get(r.context, key).Result()

	if err == redis.Nil {
		return "", errors.New("key does not exist")
	}

	if err != nil {
		return "", errors.New("err not nil attempting to Get key from Redis")
	}

	return value, nil
}

// Set a key value pair in our Redis DB
func (r *Redis) Set(key string, value string, expiration time.Duration) error {
	return r.client.Set(r.context, key, value, expiration).Err()
}

func (r *Redis) GetRawConnectionAndContext() (*redis.Client, context.Context) {
	return r.client, r.context
}
