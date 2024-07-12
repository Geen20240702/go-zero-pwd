package logic

import (
	"context"

	"go-zero-pwd/app/pwd/rpc/internal/svc"
	"go-zero-pwd/app/pwd/rpc/pwd"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPwdLogic {
	return &DelPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelPwdLogic) DelPwd(in *pwd.DelPwdReq) (*pwd.DelPwdResp, error) {
	// todo: add your logic here and delete this line

	return &pwd.DelPwdResp{}, nil
}
