package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/ljp-lachouchou/chan_xin/pkg/ctxdata"
	"github.com/ljp-lachouchou/chan_xin/pkg/lerr"
	"github.com/pkg/errors"
	"time"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	ErrNoPagination = errors.New("没有分页信息，无法浏览")
)

type ListVisiblePostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListVisiblePostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListVisiblePostsLogic {
	return &ListVisiblePostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 浏览可见动态流（根据权限过滤+分页）
func (l *ListVisiblePostsLogic) ListVisiblePosts(in *dynamics.ListVisiblePostsRequest) (*dynamics.PostListResponse, error) {
	if in.Pagination == nil {
		return nil, errors.WithStack(ErrNoPagination)
	}
	var startLocation int32 = 0
	if in.Pagination.PageToken == "" {
		//首次
		return l.listVisiblePosts(in.Pagination.PageSize, startLocation, in.ViewerId)
	}
	secretKey := []byte(l.svcCtx.Config.JwtAuth.AccessSecret)
	parse, err2 := ctxdata.ParseByTokenString(secretKey, in.Pagination.PageToken)
	if err2 != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err2, "dynamics-rpc ListVisiblePosts ctxdata.GenPageToken")
	}
	v, ok := parse.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.WithStack(ChangeMapClamsErr)
	}
	startLocation = int32(v[ctxdata.PageIdentify].(float64))
	return l.listVisiblePosts(in.Pagination.PageSize, startLocation, in.ViewerId)
}
func (l *ListVisiblePostsLogic) listVisiblePosts(pageSize, offset int32, userId string) (*dynamics.PostListResponse, error) {
	endLocation := offset + pageSize
	posts, err := l.svcCtx.PostsModel.FindCanVisiablePosts(l.ctx, userId, int64(offset), int64(pageSize))
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NEWDBError(), err, "FindCanVisiblePosts", userId, pageSize, offset)
	}
	postDynamics := copyPosts(posts)
	nextToken, err := ctxdata.GenPageToken(l.svcCtx.Config.JwtAuth.AccessSecret, time.Now().Unix(), l.svcCtx.Config.JwtAuth.AccessExpire, endLocation)
	if err != nil {
		return nil, lerr.NewWrapError(lerr.NewCOMMONError(), err, "ListVisiblePosts GenPageToken", endLocation)
	}
	return &dynamics.PostListResponse{
		Posts:         postDynamics,
		NextPageToken: nextToken,
	}, nil
}
