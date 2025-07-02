package lerr

var errMapping map[ErrType]string = map[ErrType]string{
	SERVICE_COMMON_ERROR: "服务异常，请稍后再试",
	DB_ERROR:             "数据库异常，请重试",
	SYSTEM_ERROR:         "请检查错误",
}

func ErrMsg(code ErrType) string {
	if v, ok := errMapping[code]; ok {
		return v
	}
	return errMapping[SYSTEM_ERROR]
}
