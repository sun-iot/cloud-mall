package customer

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CustomerBalanceLogModel = (*customCustomerBalanceLogModel)(nil)

type (
	// CustomerBalanceLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerBalanceLogModel.
	CustomerBalanceLogModel interface {
		customerBalanceLogModel
	}

	customCustomerBalanceLogModel struct {
		*defaultCustomerBalanceLogModel
	}
)

// NewCustomerBalanceLogModel returns a model for the database table.
func NewCustomerBalanceLogModel(conn sqlx.SqlConn) CustomerBalanceLogModel {
	return &customCustomerBalanceLogModel{
		defaultCustomerBalanceLogModel: newCustomerBalanceLogModel(conn),
	}
}
