package lerr

import "github.com/zeromicro/x/errors"

func NewError(code int, message string) error {
	return errors.New(code, message)
}
func NEWDBError() error {
	return errors.New(int(DB_ERROR), ErrMsg(DB_ERROR))
}
func NewCOMMONError() error {
	return errors.New(int(SERVICE_COMMON_ERROR), ErrMsg(SERVICE_COMMON_ERROR))
}
func NewSYSTEMError() error {
	return errors.New(int(SYSTEM_ERROR), ErrMsg(SYSTEM_ERROR))
}
