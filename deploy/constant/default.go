package constant

type FriendApplyHandle int64

const (
	NoHandleApply FriendApplyHandle = iota
	SuccessHandleApply
	FailHandleApply
)

type IsAdminInGroup int64

const (
	NoAdminInGroup IsAdminInGroup = iota
	AdminInGroup
)
