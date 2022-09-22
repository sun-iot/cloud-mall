package customer

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CustomerLoginModel = (*customCustomerLoginModel)(nil)

type (
	// CustomerLoginModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerLoginModel.
	CustomerLoginModel interface {
		customerLoginModel
	}

	customCustomerLoginModel struct {
		*defaultCustomerLoginModel
	}
)

// NewCustomerLoginModel returns a model for the database table.
func NewCustomerLoginModel(conn sqlx.SqlConn) CustomerLoginModel {
	return &customCustomerLoginModel{
		defaultCustomerLoginModel: newCustomerLoginModel(conn),
	}
}
