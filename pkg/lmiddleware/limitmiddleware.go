package lmiddleware

import (
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

type LimitMiddleware struct {
	redisCfg redis.RedisConf
	*limit.TokenLimiter
}

func NewLimitMiddleware(redisCfg redis.RedisConf) *LimitMiddleware {
	return &LimitMiddleware{
		redisCfg: redisCfg,
	}
}
func (m *LimitMiddleware) TokenLimitHandler(rate, burst int) rest.Middleware {
	m.TokenLimiter = limit.NewTokenLimiter(rate, burst, redis.MustNewRedis(m.redisCfg), "REDIS_TOKEN_LIMIT_KEY")

	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if m.TokenLimiter.AllowCtx(request.Context()) {
				next(writer, request)
				return
			}
			writer.WriteHeader(http.StatusUnauthorized)
		}
	}
}
