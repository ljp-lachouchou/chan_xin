// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type BaseUserInfo struct {
	UserId    string `json:"user_id"`    // 用户全局ID
	Nickname  string `json:"nickname"`   // 昵称
	AvatarUrl string `json:"avatar_url"` // 头像URL
	Gender    uint32 `json:"gender"`     // 性别 (0-未知 1-男 2-女)
}

type FriendApplyAction struct {
	ApplyId    string `json:"apply_id"`    // 申请记录ID
	IsApproved bool   `json:"is_approved"` // 是否同意
}

type FriendApplyListReq struct {
	UserId string `form:"user_id"` //用户id
}

type FriendApplyListResp struct {
	List []*FriendApplyResp `json:"list"` //用户请求添加的好友
}

type FriendApplyRequest struct {
	ApplicantId string `json:"applicant_id"` // 申请人ID
	TargetId    string `json:"target_id"`    // 目标用户ID
	GreetMsg    string `json:"greet_msg"`    // 打招呼消息
}

type FriendApplyResp struct {
	UserId    string `json:"user_id"`    // 用户全局ID
	Nickname  string `json:"nickname"`   // 昵称
	AvatarUrl string `json:"avatar_url"` // 头像URL
	Gender    uint32 `json:"gender"`     // 性别 (0-未知 1-男 2-女)
	GreetMsg  string `json:"greet_msg"`  // 打招呼消息
	Status    int32  `json:"status"`     // 申请状态
}

type FriendApplyResponse struct {
	ApplyId   string `json:"apply_id"` // 申请记录ID
	ApplyTime int64  `json:"apply_time"`
}

type FriendInfoRequest struct {
	UserId   string `form:"user_id"`   // 查询者ID
	FriendId string `form:"friend_id"` // 目标好友ID
}

type FriendListReq struct {
	UserId string `form:"user_id"` //用户id
}

type FriendListResp struct {
	FriendList []*UserInfo `json:"friend_list"` //该用户的好友
}

type FriendStatus struct {
	IsMuted   *bool   `json:"is_muted,omitempty"`   // 消息免打扰
	IsTopped  *bool   `json:"is_topped,omitempty"`  // 置顶聊天
	IsBlocked *bool   `json:"is_blocked,omitempty"` // 拉黑好友
	Remark    *string `json:"remark,omitempty"`     // 好友备注
}

type FriendStatusInfo struct {
	IsMuted   bool   `json:"is_muted"`   // 消息免打扰
	IsTopped  bool   `json:"is_topped"`  // 置顶聊天
	IsBlocked bool   `json:"is_blocked"` // 拉黑好友
	Remark    string `json:"remark"`     // 好友备注
}

type FriendStatusUpdate struct {
	UserId   string       `json:"user_id"`   // 操作者ID
	FriendId string       `json:"friend_id"` // 好友ID
	Status   FriendStatus `json:"status"`    // 新状态
}

type GetGroupMembersReq struct {
	GroupId string `form:"group_id"` // 群组id
}

type GetGroupMembersResp struct {
	List []*BaseUserInfo `json:"list"` //群成员的信息
}

type GroupApplyAction struct {
	ApplyId    string `json:"apply_id"`    //申请id
	ManagerId  string `json:"manager_id"`  //处理人id
	IsApproved bool   `json:"is_approved"` //是否同意申请
}

type GroupApplyReq struct {
	ApplicantId string `json:"applicant_id"` //申请方---个人
	TargetId    string `json:"target_id"`    //接收方 --- 群
	GreetMsg    string `json:"greet_msg"`    // 招呼
}

type GroupApplyResp struct {
	ApplyId   string `json:"apply_id"` //申请iD
	ApplyTime int64  `json:"apply_time"`
}

type GroupCreationRequest struct {
	CreatorId string   `json:"creator_id"` // 创建者ID
	GroupName string   `json:"group_name"` // 初始群名
	MemberIds []string `json:"member_ids"` // 初始成员ID
}

type GroupInfo struct {
	GroupId    string   `json:"group_id"`    // 群ID
	Name       string   `json:"name"`        // 群名称
	OwnerId    string   `json:"owner_id"`    // 群主ID
	AdminIds   []string `json:"admin_ids"`   // 管理员ID列表
	MemberIds  []string `json:"member_ids"`  // 成员ID列表
	MaxMembers uint32   `json:"max_members"` // 群人数上限
}

type GroupInfoRequest struct {
	UserId  string `form:"user_id"`  // 查询者ID
	GroupId string `form:"group_id"` // 目标群ID
}

type GroupInvitation struct {
	InviterId string   `json:"inviter_id"` // 邀请人ID
	GroupId   string   `json:"group_id"`   // 群ID
	TargetIds []string `json:"target_ids"` // 被邀请人ID列表
}

type GroupInviteAction struct {
	InviteId   string `json:"invite_id"`   // 邀请记录ID
	IsAccepted bool   `json:"is_accepted"` // 是否接受
}

type GroupMemberManage struct {
	OperatorId string          `json:"operator_id"` // 操作者ID
	GroupId    string          `json:"group_id"`
	TargetId   string          `json:"target_id"` // 被操作成员ID
	Action     GroupPermission `json:"action"`    // 操作
}

type GroupMemberSetting struct {
	GroupNickname      string `json:"group_nickname,omitempty"`       // 群内昵称
	ShowMemberNickname bool   `json:"show_member_nickname,omitempty"` // 是否显示群成员昵称
}

type GroupMemberSettingUpdate struct {
	UserId  string             `json:"user_id"`  //用户id
	GroupId string             `json:"group_id"` //群组id
	Setting GroupMemberSetting `json:"setting"`  //个性化设置
}

type GroupPermission struct {
	Action int32 `json:"action"` //0-踢出成员 1-设置/取消管理员 2-修改群名
}

type GroupQuitRequest struct {
	UserId  string `json:"user_id"`  // 退出者ID
	GroupId string `json:"group_id"` //群id
}

type GroupStatus struct {
	IsMuted  bool   `json:"is_muted,omitempty"`  // 群消息免打扰
	IsTopped bool   `json:"is_topped,omitempty"` // 置顶群聊
	Remark   string `json:"remark,omitempty"`    // 群备注
}

type GroupStatusUpdate struct {
	UserId  string      `json:"user_id"`  // 操作者ID
	GroupId string      `json:"group_id"` //群id
	Status  GroupStatus `json:"status"`   // 新状态
}

type RelationRequest struct {
	FromUid string `json:"from_uid"` // 发起方
	ToUid   string `json:"to_uid"`   // 目标方
}

type RemoveAdminReq struct {
	OperatorId string `json:"operator_id"` // 操作者ID
	GroupId    string `json:"group_id"`
	TargetId   string `json:"target_id"` // 被操作成员ID
}

type UserInfo struct {
	UserId    string           `json:"user_id"`    // 用户全局ID
	Nickname  string           `json:"nickname"`   // 昵称
	AvatarUrl string           `json:"avatar_url"` // 头像URL
	Gender    uint32           `json:"gender"`     // 性别 (0-未知 1-男 2-女)
	Status    FriendStatusInfo `json:"status"`     // 对于好友的状态
}
