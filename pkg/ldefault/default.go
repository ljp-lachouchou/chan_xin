package ldefault

import "time"

const (
	SYSTEM_REDIS_TOPNKEN_KEY                = "root:token"
	DEFAULT_EXP                             = 86400 * 2
	SYSTEM_REDIS_UID                        = "root"
	DEFAULT_MAX_GROUP_MEMBERS               = 200
	DEFAULT_REDIS_LOCK_KEY                  = "redis_lock"
	DEFAULT_REDIS_LOCK_EXPIRE time.Duration = 20 * time.Second
)
