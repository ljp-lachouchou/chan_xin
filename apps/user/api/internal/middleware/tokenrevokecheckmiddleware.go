package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ljp-lachouchou/chan_xin/apps/user/api/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"net/http"
	"strings"
)

type TokenRevokeCheckMiddleware struct {
}

func NewTokenRevokeCheckMiddleware() *TokenRevokeCheckMiddleware {
	return &TokenRevokeCheckMiddleware{}
}

func (m *TokenRevokeCheckMiddleware) Handle(svc *svc.ServiceContext) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims := &jwt.MapClaims{}
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				return []byte(svc.Config.JwtAuth.AccessSecret), nil
			}
			parse, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)

			// 2. 拦截解析失败
			if err != nil || !parse.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("令牌无效或已过期"))
				return
			}

			// 3. 提取UID并查询Redis
			uid := (*claims)[ctxdata.Identify].(string) // 根据实际声明字段调整
			redisKey := fmt.Sprintf("user_token:%s", uid)
			fmt.Println("key ", redisKey)
			latestToken, err := svc.Redis.Get(redisKey)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("系统错误"))
				return
			}
			fmt.Println("new ", latestToken, "\n old ", tokenString)
			if latestToken == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("登录验证过期,请重新进行登录"))
				return
			}
			// 4. 校验令牌是否最新
			if latestToken != "" && latestToken != tokenString {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("账号已在其他设备登录"))
				return
			}

			// 5. 验证通过，执行后续逻辑
			next(w, r)
		}
	}
}
