// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package socialmodels

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupOperationFieldNames          = builder.RawFieldNames(&GroupOperation{})
	groupOperationRows                = strings.Join(groupOperationFieldNames, ",")
	groupOperationRowsExpectAutoSet   = strings.Join(stringx.Remove(groupOperationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupOperationRowsWithPlaceHolder = strings.Join(stringx.Remove(groupOperationFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGroupOperationIdPrefix = "cache:groupOperation:id:"
	cacheGroupOperationPrefix = "cache:groupOperation:gid:oid:tid:"
)

type (
	groupOperationModel interface {
		Insert(ctx context.Context, data *GroupOperation) (sql.Result, error)
		InertWithSession(ctx context.Context,session sqlx.Session, data *GroupOperation) (sql.Result, error)
		FindOne(ctx context.Context, id uint64) (*GroupOperation, error)
		Update(ctx context.Context, data *GroupOperation) error
		Delete(ctx context.Context, id uint64) error
		DeleteByGIdOidTid(ctx context.Context,session sqlx.Session,gid,oid,tid string) error
		Transx(ctx context.Context,fn func(ctx context.Context, session sqlx.Session) error) error
	}

	defaultGroupOperationModel struct {
		sqlc.CachedConn
		table string
	}

	GroupOperation struct {
		Id         uint64         `db:"id"`
		GroupId    string         `db:"group_id"`    // 群ID
		OperatorId string         `db:"operator_id"` // 操作者ID
		TargetId   sql.NullString `db:"target_id"`   // 被操作成员ID
		ActionType string         `db:"action_type"` // 操作类型
		ExtraInfo  sql.NullString `db:"extra_info"`  // 扩展信息 (如新群名)
		CreatedAt  time.Time      `db:"created_at"`  // 操作时间
	}
)

func newGroupOperationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGroupOperationModel {
	return &defaultGroupOperationModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`group_operation`",
	}
}
func (m *defaultGroupOperationModel) DeleteByGIdOidTid(ctx context.Context,session sqlx.Session,gid,oid,tid string) error {
	groupOperationKey := fmt.Sprintf("%s%v:%v:%v", cacheGroupOperationPrefix, gid,oid,tid)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `group_id` = ? and `operator_id` = ? and `target_id` = ?", m.table)
		return session.ExecCtx(ctx, query, gid, oid, tid)
	}, groupOperationKey)
	return err

}
func (m *defaultGroupOperationModel) Transx(ctx context.Context,fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}
func (m *defaultGroupOperationModel) Delete(ctx context.Context, id uint64) error {
	groupOperationIdKey := fmt.Sprintf("%s%v", cacheGroupOperationIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, groupOperationIdKey)
	return err
}

func (m *defaultGroupOperationModel) FindOne(ctx context.Context, id uint64) (*GroupOperation, error) {
	groupOperationIdKey := fmt.Sprintf("%s%v", cacheGroupOperationIdPrefix, id)
	var resp GroupOperation
	err := m.QueryRowCtx(ctx, &resp, groupOperationIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupOperationRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupOperationModel) Insert(ctx context.Context, data *GroupOperation) (sql.Result, error) {
	groupOperationKey := fmt.Sprintf("%s%v:%v:%v", cacheGroupOperationPrefix, data.GroupId,data.OperatorId,data.TargetId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupOperationRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.GroupId, data.OperatorId, data.TargetId, data.ActionType, data.ExtraInfo)
	}, groupOperationKey)
	return ret, err
}
func (m *defaultGroupOperationModel) InertWithSession(ctx context.Context,session sqlx.Session, data *GroupOperation) (sql.Result, error) {
	groupOperationKey := fmt.Sprintf("%s%v:%v:%v", cacheGroupOperationPrefix, data.GroupId,data.OperatorId,data.TargetId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupOperationRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.GroupId, data.OperatorId, data.TargetId, data.ActionType, data.ExtraInfo)
	}, groupOperationKey)
	return ret, err
}

func (m *defaultGroupOperationModel) Update(ctx context.Context, data *GroupOperation) error {
	groupOperationIdKey := fmt.Sprintf("%s%v", cacheGroupOperationIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupOperationRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.GroupId, data.OperatorId, data.TargetId, data.ActionType, data.ExtraInfo, data.Id)
	}, groupOperationIdKey)
	return err
}

func (m *defaultGroupOperationModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGroupOperationIdPrefix, primary)
}

func (m *defaultGroupOperationModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupOperationRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupOperationModel) tableName() string {
	return m.table
}
