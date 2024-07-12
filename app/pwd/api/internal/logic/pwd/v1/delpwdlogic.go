package v1

import (
	"context"

	"go-zero-pwd/app/pwd/api/internal/svc"
	"go-zero-pwd/app/pwd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPwdLogic {
	return &DelPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelPwdLogic) DelPwd(req *types.DelPwdReq) (resp *types.DelPwdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
