package redisUtil

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type Redis struct {
	Client *redis.Client
}
type Opt struct {
	Addr     string
	Password string
	DB       int
}

func NewRedis(opt *Opt) *Redis {
	return &Redis{redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       0,
	})}
}
func (r *Redis) Set(key string, val any, expiration time.Duration, prefix ...any) error {
	return r.Client.Set(context.Background(), r.JoinKey(key, prefix...), val, expiration).Err()
}
func (r *Redis) Get(key string, prefix ...any) string {
	return r.Client.Get(context.Background(), r.JoinKey(key, prefix...)).String()
}
func (r *Redis) Del(key string, prefix ...any) error {
	return r.Client.Del(context.Background(), r.JoinKey(key, prefix...)).Err()
}

// SAdd 添加至合集
func (r *Redis) SAdd(key string, val ...any) error {
	return r.Client.SAdd(context.Background(), key, val...).Err()
}

// SRem 在合集中移除
func (r *Redis) SRem(key string, val ...any) error {
	return r.Client.SRem(context.Background(), key, val...).Err()
}

// SExists 在合集中查询值是否存在
func (r *Redis) SExists(key string, val any, prefix ...any) bool {
	ok, _ := r.Client.SIsMember(context.Background(), r.JoinKey(key, prefix...), val).Result()
	return ok
}

// JoinKey Key添加前缀
func (r *Redis) JoinKey(key string, prefix ...any) string {
	if len(prefix) == 0 {
		return key
	}
	return fmt.Sprintf("%s-%v", key, prefix[0])
}
