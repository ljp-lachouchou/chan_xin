package lerr

type ErrType int

const (
	DB_ERROR ErrType = iota + 10000
	SERVICE_COMMON_ERROR
	SYSTEM_ERROR
)
