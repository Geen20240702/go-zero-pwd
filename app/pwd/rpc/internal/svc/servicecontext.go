package svc

import (
	"go-zero-pwd/app/pwd/rpc/internal/config"
	"go-zero-pwd/app/pwd/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	PwdModel model.PwdModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		PwdModel: model.NewPwdModel(sqlx.NewMysql(c.Datasource), c.Cache),
	}
}
