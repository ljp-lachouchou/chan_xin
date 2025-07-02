package redislock

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"time"
)

type DistributedLock interface {
	Acquire() (bool, error)
	Release() (bool, error)
}
type redisLock struct {
	DistributedLock
	*redis.Redis
	Key    string
	Value  string
	Expire time.Duration
}

func NewRedisLock(redis *redis.Redis, key string, value string, expire time.Duration) DistributedLock {
	return &redisLock{Redis: redis, Key: key, Value: value, Expire: expire}
}
func (r *redisLock) Acquire() (bool, error) {
	err := r.Redis.Setex(r.Key, r.Value, int(r.Expire.Seconds()))
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *redisLock) Release() (bool, error) {
	script := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else 
			return 0
		end
	`
	eval, err := r.Redis.Eval(script, []string{r.Key}, []string{r.Value})
	if err != nil {
		return false, err
	}
	return eval.(int64) == 1, nil

}
