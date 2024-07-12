package svc

import (
	"go-zero-pwd/app/pwd/api/internal/config"
	"go-zero-pwd/app/pwd/rpc/pwdclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	PwdRpc pwdclient.Pwd
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PwdRpc: pwdclient.NewPwd(zrpc.MustNewClient(c.PwdRpcConf)),
	}
}
