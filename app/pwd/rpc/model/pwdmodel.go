package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ PwdModel = (*customPwdModel)(nil)

type (
	// PwdModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPwdModel.
	PwdModel interface {
		pwdModel
	}

	customPwdModel struct {
		*defaultPwdModel
	}
)

// NewPwdModel returns a model for the database table.
func NewPwdModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PwdModel {
	return &customPwdModel{
		defaultPwdModel: newPwdModel(conn, c, opts...),
	}
}
