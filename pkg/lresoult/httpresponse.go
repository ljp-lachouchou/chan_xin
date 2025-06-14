package lresoult

import (
	"context"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	zrpcErr "github.com/zeromicro/x/errors"
	"google.golang.org/grpc/status"
	"net/http"
)

type response struct {
	code int         `json:"code"`
	msg  string      `json:"msg"`
	data interface{} `json:"data"`
}

func Success(data interface{}) *response {
	return &response{code: 200, msg: "success", data: data}
}
func Fail(code int, msg string) *response {
	return &response{code: code, msg: msg, data: nil}
}
func OkHandler(_ context.Context, data any) any {
	return Success(data)
}
func ErrorHandler(name string) func(err error) (int, any) {
	return func(err error) (int, any) {
		errCode := lerr.SERVICE_COMMON_ERROR
		errMsg := lerr.ErrMsg(lerr.SERVICE_COMMON_ERROR)
		causeErr := errors.Cause(err)

		if msg, ok := causeErr.(*zrpcErr.CodeMsg); ok { //zero
			errCode = lerr.ErrType(msg.Code)
			errMsg = msg.Msg
		} else {
			if gstatus, ok := status.FromError(err); ok { //grpc错误状态1
				errCode = lerr.ErrType(gstatus.Code())
				errMsg = gstatus.Message()
			}
		}
		//日志记录
		logx.WithCallerSkip(0).Errorf("【%s】 err %v", name, err)
		return http.StatusBadRequest, Fail(int(errCode), errMsg)
	}
}
