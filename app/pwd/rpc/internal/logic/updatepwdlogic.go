package logic

import (
	"context"

	"go-zero-pwd/app/pwd/rpc/internal/svc"
	"go-zero-pwd/app/pwd/rpc/pwd"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePwdLogic {
	return &UpdatePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePwdLogic) UpdatePwd(in *pwd.UpdatePwdReq) (*pwd.UpdatePwdResp, error) {
	// todo: add your logic here and delete this line

	return &pwd.UpdatePwdResp{}, nil
}
