package logic

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/dynamicsmodels"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"
	"time"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	PaginationNilErr  = errors.New("页面配置查询不到")
	ChangeMapClamsErr = errors.New("转换为MapClaims失败")
)

type ListUserPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUserPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUserPostsLogic {
	return &ListUserPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户所有动态（按置顶状态+时间倒序）
func (l *ListUserPostsLogic) ListUserPosts(in *dynamics.ListUserPostsRequest) (*dynamics.PostListResponse, error) {
	if in.IsPin != nil && *in.IsPin {
		if in.Pagination == nil {
			return nil, errors.WithStack(PaginationNilErr)
		}
		var startLocation int32
		startLocation = 0
		if in.Pagination.PageToken == "" {
			return l.listUserByPosts(true, in.Pagination.PageSize, startLocation, in.UserId, fmt.Sprintf("start location pin token nil: %d", startLocation))
		}
		secretKey := []byte(l.svcCtx.Config.JwtAuth.AccessSecret)
		parse, err2 := ctxdata.ParseByTokenString(secretKey, in.Pagination.PageToken)
		if err2 != nil {
			return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err2, "dynamics-rpc ListUserPosts ctxdata.GenPageToken 1")
		}
		v, ok := parse.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.WithStack(ChangeMapClamsErr)
		}
		startLocation = int32(v[ctxdata.PageIdentify].(float64))

		return l.listUserByPosts(true, in.Pagination.PageSize, startLocation, in.UserId, fmt.Sprintf("start location pin not token nil: %d", startLocation))
	}
	if in.Pagination == nil {
		return nil, errors.WithStack(PaginationNilErr)
	}
	var startLocation int32 = 0
	if in.Pagination.PageToken == "" {
		return l.listUserByPosts(false, in.Pagination.PageSize, startLocation, in.UserId, fmt.Sprintf("start location not pin token nil: %d", startLocation))
	}
	secretKey := []byte(l.svcCtx.Config.JwtAuth.AccessSecret)
	parse, err2 := ctxdata.ParseByTokenString(secretKey, in.Pagination.PageToken)
	if err2 != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err2, "dynamics-rpc ListUserPosts ctxdata.GenPageToken 1")
	}
	v, ok := parse.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.WithStack(ChangeMapClamsErr)
	}
	startLocation = int32(v[ctxdata.PageIdentify].(float64))
	return l.listUserByPosts(false, in.Pagination.PageSize, startLocation, in.UserId, fmt.Sprintf("start location not pin not token nil: %d", startLocation))
}
func copyPosts(posts []*dynamicsmodels.Posts) []*dynamics.Post {
	var postsDynamics []*dynamics.Post
	for _, v := range posts {
		post := &dynamics.Post{
			PostId: v.ID.Hex(),
			UserId: v.UserId,
			Content: &dynamics.PostContent{
				Text:      v.Content,
				ImageUrls: v.ImageUrls,
				Emoji:     v.Emoji,
			},
			Meta: &dynamics.PostMeta{
				Location:       v.Location,
				Scope:          dynamics.VisibleScope(v.Visibility),
				VisibleUserIds: v.VisibleTo,
			},
			IsPinned: v.IsPinned,
			CreateTime: v.CreateTime,
		}
		postsDynamics = append(postsDynamics, post)
	}
	return postsDynamics
}
func (l *ListUserPostsLogic) listUserByPosts(isPin bool, pageSize, startLocation int32, UserId, fString string) (*dynamics.PostListResponse, error) {
	endLocation := pageSize + startLocation
	l.Info(fString)
	token, err := ctxdata.GenPageToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire, endLocation)
	nextToken := token
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "dynamics-rpc ListUserPosts ctxdata.GenPageToken")
	}
	posts, err := l.svcCtx.PostsModel.FindPostsByUserId(l.ctx, UserId, int64(startLocation), int64(pageSize), isPin)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "dynamics-rpc ListUserPosts FindPostsByUserId")
	}
	postsDynamics := copyPosts(posts)
	return &dynamics.PostListResponse{
		Posts:         postsDynamics,
		NextPageToken: nextToken,
	}, nil
}
