package handler

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ljp-lachouchou/chan_xin/apps/im/ws/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/token"
	"net/http"
)

type Auth struct {
	svc   *svc.ServiceContext
	parse *token.TokenParser
	logx.Logger
}

func NewAuth(svc *svc.ServiceContext) *Auth {
	return &Auth{
		svc:    svc,
		parse:  token.NewTokenParser(),
		Logger: logx.WithContext(context.Background()),
	}
}
func (a *Auth) Auth(w http.ResponseWriter, r *http.Request) bool {
	get := r.Header.Get("Authorization")
	fmt.Println("Authorization==", get)
	parseToken, err := a.parse.ParseToken(r, a.svc.Jwt.AccessSecret, "")
	if err != nil {
		a.Errorf("parse token err: %v", err)
		parseToken, err = a.parse.ParseToken(r, a.svc.Jwt.AccessSecret, "Brearer")
		if err != nil {
			return false
		}
	}
	if !parseToken.Valid {
		a.Errorf("token is invalid")
		return false
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	*r = *r.WithContext(context.WithValue(r.Context(), ctxdata.Identify, claims[ctxdata.Identify]))
	return true
}
func (a *Auth) GetUid(r *http.Request) string {
	return ctxdata.GetUId(r.Context())
}
