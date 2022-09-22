package customer

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CustomerPointLogModel = (*customCustomerPointLogModel)(nil)

type (
	// CustomerPointLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerPointLogModel.
	CustomerPointLogModel interface {
		customerPointLogModel
	}

	customCustomerPointLogModel struct {
		*defaultCustomerPointLogModel
	}
)

// NewCustomerPointLogModel returns a model for the database table.
func NewCustomerPointLogModel(conn sqlx.SqlConn) CustomerPointLogModel {
	return &customCustomerPointLogModel{
		defaultCustomerPointLogModel: newCustomerPointLogModel(conn),
	}
}
