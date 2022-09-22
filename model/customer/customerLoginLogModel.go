package customer

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CustomerLoginLogModel = (*customCustomerLoginLogModel)(nil)

type (
	// CustomerLoginLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerLoginLogModel.
	CustomerLoginLogModel interface {
		customerLoginLogModel
	}

	customCustomerLoginLogModel struct {
		*defaultCustomerLoginLogModel
	}
)

// NewCustomerLoginLogModel returns a model for the database table.
func NewCustomerLoginLogModel(conn sqlx.SqlConn) CustomerLoginLogModel {
	return &customCustomerLoginLogModel{
		defaultCustomerLoginLogModel: newCustomerLoginLogModel(conn),
	}
}
