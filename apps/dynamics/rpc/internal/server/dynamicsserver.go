// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: dynamics.proto

package server

import (
	"context"

	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/dynamics"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/logic"
	"github.com/ljp-lachouchou/chan_xin/apps/dynamics/rpc/internal/svc"
)

type DynamicsServer struct {
	svcCtx *svc.ServiceContext
	dynamics.UnimplementedDynamicsServer
}

func NewDynamicsServer(svcCtx *svc.ServiceContext) *DynamicsServer {
	return &DynamicsServer{
		svcCtx: svcCtx,
	}
}

// 创建动态（需提供内容和隐私设置）
func (s *DynamicsServer) CreatePost(ctx context.Context, in *dynamics.CreatePostRequest) (*dynamics.Post, error) {
	l := logic.NewCreatePostLogic(ctx, s.svcCtx)
	return l.CreatePost(in)
}

// 删除动态（仅创建者可操作）
func (s *DynamicsServer) DeletePost(ctx context.Context, in *dynamics.DeletePostRequest) (*dynamics.Empty, error) {
	l := logic.NewDeletePostLogic(ctx, s.svcCtx)
	return l.DeletePost(in)
}

// 点赞/取消点赞
func (s *DynamicsServer) ToggleLike(ctx context.Context, in *dynamics.LikeAction) (*dynamics.Empty, error) {
	l := logic.NewToggleLikeLogic(ctx, s.svcCtx)
	return l.ToggleLike(in)
}

// 置顶/取消置顶动态
func (s *DynamicsServer) PinPost(ctx context.Context, in *dynamics.PinPostRequest) (*dynamics.Empty, error) {
	l := logic.NewPinPostLogic(ctx, s.svcCtx)
	return l.PinPost(in)
}

// 获取用户所有动态（按置顶状态+时间倒序）
func (s *DynamicsServer) ListUserPosts(ctx context.Context, in *dynamics.ListUserPostsRequest) (*dynamics.PostListResponse, error) {
	l := logic.NewListUserPostsLogic(ctx, s.svcCtx)
	return l.ListUserPosts(in)
}

// 设置个人动态封面（用于个人主页）
func (s *DynamicsServer) SetCover(ctx context.Context, in *dynamics.SetCoverRequest) (*dynamics.Empty, error) {
	l := logic.NewSetCoverLogic(ctx, s.svcCtx)
	return l.SetCover(in)
}

// 浏览可见动态流（根据权限过滤+分页）
func (s *DynamicsServer) ListVisiblePosts(ctx context.Context, in *dynamics.ListVisiblePostsRequest) (*dynamics.PostListResponse, error) {
	l := logic.NewListVisiblePostsLogic(ctx, s.svcCtx)
	return l.ListVisiblePosts(in)
}

// 创建评论
func (s *DynamicsServer) CreateComment(ctx context.Context, in *dynamics.CreateCommentReq) (*dynamics.CreateCommentResp, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

// 创建评论回复
func (s *DynamicsServer) CreateCommentReplay(ctx context.Context, in *dynamics.CreateCommentReplayReq) (*dynamics.CreateCommentReplayResp, error) {
	l := logic.NewCreateCommentReplayLogic(ctx, s.svcCtx)
	return l.CreateCommentReplay(in)
}

// 更新评论
func (s *DynamicsServer) UpdateComment(ctx context.Context, in *dynamics.UpdateCommentReq) (*dynamics.Empty, error) {
	l := logic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

// 更新评论回复
func (s *DynamicsServer) UpdateCommentReplay(ctx context.Context, in *dynamics.UpdateCommentReplayReq) (*dynamics.Empty, error) {
	l := logic.NewUpdateCommentReplayLogic(ctx, s.svcCtx)
	return l.UpdateCommentReplay(in)
}

// 更新通知
func (s *DynamicsServer) UpdateNotification(ctx context.Context, in *dynamics.UpdateNotificationReq) (*dynamics.Empty, error) {
	l := logic.NewUpdateNotificationLogic(ctx, s.svcCtx)
	return l.UpdateNotification(in)
}

// 删除评论
func (s *DynamicsServer) DeleteComment(ctx context.Context, in *dynamics.DeleteCommentReq) (*dynamics.Empty, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

// 删除评论回复
func (s *DynamicsServer) DeleteCommentReplay(ctx context.Context, in *dynamics.DeleteCommentReplayReq) (*dynamics.Empty, error) {
	l := logic.NewDeleteCommentReplayLogic(ctx, s.svcCtx)
	return l.DeleteCommentReplay(in)
}

// 创建通知
func (s *DynamicsServer) CreateNotification(ctx context.Context, in *dynamics.CreateNotificationReq) (*dynamics.Empty, error) {
	l := logic.NewCreateNotificationLogic(ctx, s.svcCtx)
	return l.CreateNotification(in)
}

// 获取通知列表（分页）
func (s *DynamicsServer) ListNotifications(ctx context.Context, in *dynamics.ListNotificationsRequest) (*dynamics.ListNotificationsResponse, error) {
	l := logic.NewListNotificationsLogic(ctx, s.svcCtx)
	return l.ListNotifications(in)
}

// 单个post信息
func (s *DynamicsServer) GetPostInfo(ctx context.Context, in *dynamics.GetPostInfoReq) (*dynamics.Post, error) {
	l := logic.NewGetPostInfoLogic(ctx, s.svcCtx)
	return l.GetPostInfo(in)
}

// 点赞列表
func (s *DynamicsServer) ListLikeByPostId(ctx context.Context, in *dynamics.GetPostInfoReq) (*dynamics.Ids, error) {
	l := logic.NewListLikeByPostIdLogic(ctx, s.svcCtx)
	return l.ListLikeByPostId(in)
}

// 评论列表
func (s *DynamicsServer) ListCommentByPostId(ctx context.Context, in *dynamics.GetPostInfoReq) (*dynamics.ListCommentResp, error) {
	l := logic.NewListCommentByPostIdLogic(ctx, s.svcCtx)
	return l.ListCommentByPostId(in)
}

// 根据type和userid查找
func (s *DynamicsServer) ListNotificationsByUserIdAndType(ctx context.Context, in *dynamics.ListNotificationsByUserIdAndTypeReq) (*dynamics.ListNotificationsByUserIdAndTypeReqResponse, error) {
	l := logic.NewListNotificationsByUserIdAndTypeLogic(ctx, s.svcCtx)
	return l.ListNotificationsByUserIdAndType(in)
}

// 新增：获取未读通知数量
func (s *DynamicsServer) GetUnreadCount(ctx context.Context, in *dynamics.GetUnreadCountRequest) (*dynamics.GetUnreadCountResponse, error) {
	l := logic.NewGetUnreadCountLogic(ctx, s.svcCtx)
	return l.GetUnreadCount(in)
}

func (s *DynamicsServer) Ping(ctx context.Context, in *dynamics.PingRep) (*dynamics.PingResp, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
