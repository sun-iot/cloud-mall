package customer

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CustomerLevelModel = (*customCustomerLevelModel)(nil)

type (
	// CustomerLevelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerLevelModel.
	CustomerLevelModel interface {
		customerLevelModel
	}

	customCustomerLevelModel struct {
		*defaultCustomerLevelModel
	}
)

// NewCustomerLevelModel returns a model for the database table.
func NewCustomerLevelModel(conn sqlx.SqlConn) CustomerLevelModel {
	return &customCustomerLevelModel{
		defaultCustomerLevelModel: newCustomerLevelModel(conn),
	}
}
