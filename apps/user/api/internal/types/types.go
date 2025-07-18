// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package types

type FindUserReq struct {
	Name  string   `form:"name",omitempty`  //昵称
	Phone string   `form:"phone",omitempty` //电话
	Ids   []string `form:"ids",omitempty`   //id
}

type FindUserResp struct {
	Infos []User `json:"infos"`
}

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResp struct {
	Id     string `json:"id"`     //用户id
	Token  string `json:"token"`  //用户登录token
	Expire int64  `json:"expire"` //token有效时长
}

type RegisterReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Sex      byte   `json:"sex"`
	Avatar   string `json:"avatar"`
}

type RegisterResp struct {
	Token  string `json:"token"`  //token
	Expire int64  `json:"expire"` //token有效时长
}

type UpdateReq struct {
	Nickname *string `json:"nickname,optional"` //昵称
	Avatar   *string `json:"avatar,optional"`   // 头像
	Sex      *int32  `json:"sex,optional"`      //性别
}

type UpdateResp struct {
	Info User `json:"info"`
}

type User struct {
	Id       string `json:"id"`
	Phone    string `json:"phone""`
	Nickname string `json:"nickname"`
	Sex      byte   `json:"sex"`
	Avatar   string `json:"avatar"`
}

type UserInfoResp struct {
	Info User `json:"info"` //用户信息
}
