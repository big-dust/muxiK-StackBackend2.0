// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	evaluationInfoFieldNames          = builder.RawFieldNames(&EvaluationInfo{})
	evaluationInfoRows                = strings.Join(evaluationInfoFieldNames, ",")
	evaluationInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(evaluationInfoFieldNames, "`pid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	evaluationInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(evaluationInfoFieldNames, "`pid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	evaluationInfoModel interface {
		Insert(ctx context.Context, data *EvaluationInfo) (sql.Result, error)
		FindOne(ctx context.Context, pid int64) (*EvaluationInfo, error)
		FindOneByPid(ctx context.Context, pid int64) (*EvaluationInfo, error)
		Update(ctx context.Context, data *EvaluationInfo) error
		Delete(ctx context.Context, pid int64) error
	}

	defaultEvaluationInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	EvaluationInfo struct {
		Pid      int64         `db:"pid"`
		Sid      string        `db:"sid"` // 学号
		Cid      string        `db:"cid"` // 课程号
		Folded   int64         `db:"folded"`
		Deleted  int64         `db:"deleted"`
		Info     string        `db:"info"` // 评课内容
		Liked    sql.NullInt64 `db:"liked"`
		Disliked sql.NullInt64 `db:"disliked"`
	}
)

func newEvaluationInfoModel(conn sqlx.SqlConn) *defaultEvaluationInfoModel {
	return &defaultEvaluationInfoModel{
		conn:  conn,
		table: "`evaluation_info`",
	}
}

func (m *defaultEvaluationInfoModel) withSession(session sqlx.Session) *defaultEvaluationInfoModel {
	return &defaultEvaluationInfoModel{
		conn:  sqlx.NewSqlConnFromSession(session),
		table: "`evaluation_info`",
	}
}

func (m *defaultEvaluationInfoModel) Delete(ctx context.Context, pid int64) error {
	query := fmt.Sprintf("delete from %s where `pid` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, pid)
	return err
}

func (m *defaultEvaluationInfoModel) FindOne(ctx context.Context, pid int64) (*EvaluationInfo, error) {
	query := fmt.Sprintf("select %s from %s where `pid` = ? limit 1", evaluationInfoRows, m.table)
	var resp EvaluationInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, pid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultEvaluationInfoModel) FindOneByPid(ctx context.Context, pid int64) (*EvaluationInfo, error) {
	var resp EvaluationInfo
	query := fmt.Sprintf("select %s from %s where `pid` = ? limit 1", evaluationInfoRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, pid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultEvaluationInfoModel) Insert(ctx context.Context, data *EvaluationInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, evaluationInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Sid, data.Cid, data.Folded, data.Deleted, data.Info, data.Liked, data.Disliked)
	return ret, err
}

func (m *defaultEvaluationInfoModel) Update(ctx context.Context, newData *EvaluationInfo) error {
	query := fmt.Sprintf("update %s set %s where `pid` = ?", m.table, evaluationInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Sid, newData.Cid, newData.Folded, newData.Deleted, newData.Info, newData.Liked, newData.Disliked, newData.Pid)
	return err
}

func (m *defaultEvaluationInfoModel) tableName() string {
	return m.table
}
