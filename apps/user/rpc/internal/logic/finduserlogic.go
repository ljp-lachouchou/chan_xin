package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/ljp-lachouchou/chan_xin/apps/user/usermodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"

	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/internal/svc"
	"github.com/ljp-lachouchou/chan_xin/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	var (
		userEntitys []*usermodels.Users
		err         error
	)
	phoneUsers := make(map[string]*usermodels.Users)
	nameUsers := make(map[string]*usermodels.Users)
	idUsers := make(map[string]*usermodels.Users)

	// 分支1：手机号查询
	if in.Phone != "" {
		user, err := l.svcCtx.UsersModel.FindByPhone(l.ctx, in.Phone)
		if err == nil && user != nil {
			phoneUsers[user.Id] = user // 存到临时Map
		}
	}

	// 分支2：姓名查询
	if in.Name != "" {
		users, err := l.svcCtx.UsersModel.ListByName(l.ctx, in.Name)
		if err == nil {
			for _, u := range users {
				nameUsers[u.Id] = u // 存到临时Map
			}
		} else {
			logx.Errorf("ListByName error: %v", err)
		}
	}

	// 分支3：ID列表查询
	if len(in.Ids) > 0 {
		users, err := l.svcCtx.UsersModel.ListByIds(l.ctx, in.Ids)
		if err == nil {
			for _, u := range users {
				idUsers[u.Id] = u // 存到临时Map
			}
		} else {
			logx.Errorf("ListByIds error: %v", err)
		}
	}
	resultMap := make(map[string]*usermodels.Users)

	// 优先级：phone > name > ids
	for id, user := range phoneUsers {
		resultMap[id] = user
	}
	for id, user := range nameUsers {
		// 若手机号分支已写入，跳过覆盖
		if _, exists := resultMap[id]; !exists {
			resultMap[id] = user
		}
	}
	for id, user := range idUsers {
		if _, exists := resultMap[id]; !exists {
			resultMap[id] = user
		}
	}

	// 转换为切片
	for _, user := range resultMap {
		userEntitys = append(userEntitys, user)
	}
	if err != nil {
		return nil, errors.Wrapf(lerr.NEWDBError(), "db find err: %v req:%v ,%v, %v", err, in.Phone, in.Name, in.Ids)
	}
	var resp []*user.User
	err = copier.Copy(&resp, userEntitys)
	if err != nil {
		return nil, errors.Wrapf(lerr.NewSYSTEMError(), "copy err: %v", err)
	}
	return &user.FindUserResp{
		User: resp,
	}, nil
}
