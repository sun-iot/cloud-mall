package customer

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CustomerAddressModel = (*customCustomerAddressModel)(nil)

type (
	// CustomerAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerAddressModel.
	CustomerAddressModel interface {
		customerAddressModel
	}

	customCustomerAddressModel struct {
		*defaultCustomerAddressModel
	}
)

// NewCustomerAddressModel returns a model for the database table.
func NewCustomerAddressModel(conn sqlx.SqlConn) CustomerAddressModel {
	return &customCustomerAddressModel{
		defaultCustomerAddressModel: newCustomerAddressModel(conn),
	}
}
