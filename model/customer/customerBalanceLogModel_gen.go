// Code generated by goctl. DO NOT EDIT!

package customer

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
	customerBalanceLogFieldNames          = builder.RawFieldNames(&CustomerBalanceLog{})
	customerBalanceLogRows                = strings.Join(customerBalanceLogFieldNames, ",")
	customerBalanceLogRowsExpectAutoSet   = strings.Join(stringx.Remove(customerBalanceLogFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	customerBalanceLogRowsWithPlaceHolder = strings.Join(stringx.Remove(customerBalanceLogFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	customerBalanceLogModel interface {
		Insert(ctx context.Context, data *CustomerBalanceLog) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*CustomerBalanceLog, error)
		Update(ctx context.Context, data *CustomerBalanceLog) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCustomerBalanceLogModel struct {
		conn  sqlx.SqlConn
		table string
	}

	CustomerBalanceLog struct {
		Id          int64           `db:"id"`          // 余额日志ID
		CustomerId  int64           `db:"customer_id"` // 用户ID
		Source      sql.NullInt64   `db:"source"`      // 记录来源：1订单，2退货单
		SourceSn    sql.NullInt64   `db:"source_sn"`   // 相关单据ID
		Amount      sql.NullFloat64 `db:"amount"`      // 变动金额
		CreatedTime sql.NullTime    `db:"created_time"`
		UpdatedTime sql.NullTime    `db:"updated_time"`
		DeletedTime sql.NullTime    `db:"deleted_time"`
	}
)

func newCustomerBalanceLogModel(conn sqlx.SqlConn) *defaultCustomerBalanceLogModel {
	return &defaultCustomerBalanceLogModel{
		conn:  conn,
		table: "`customer_balance_log`",
	}
}

func (m *defaultCustomerBalanceLogModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCustomerBalanceLogModel) FindOne(ctx context.Context, id int64) (*CustomerBalanceLog, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", customerBalanceLogRows, m.table)
	var resp CustomerBalanceLog
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultCustomerBalanceLogModel) Insert(ctx context.Context, data *CustomerBalanceLog) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, customerBalanceLogRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.CustomerId, data.Source, data.SourceSn, data.Amount, data.CreatedTime, data.UpdatedTime, data.DeletedTime)
	return ret, err
}

func (m *defaultCustomerBalanceLogModel) Update(ctx context.Context, data *CustomerBalanceLog) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, customerBalanceLogRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.CustomerId, data.Source, data.SourceSn, data.Amount, data.CreatedTime, data.UpdatedTime, data.DeletedTime, data.Id)
	return err
}

func (m *defaultCustomerBalanceLogModel) tableName() string {
	return m.table
}
