package v1

import (
	"context"
	"go-zero-pwd/app/pwd/api/internal/svc"
	"go-zero-pwd/app/pwd/api/internal/types"
	"go-zero-pwd/app/pwd/rpc/pwdclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePwdLogic {
	return &CreatePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreatePwdLogic) CreatePwd(req *types.CreatePwdReq) (resp *types.CreatePwdResp, err error) {
	rpcResp, err := l.svcCtx.PwdRpc.CreatePwd(l.ctx, &pwdclient.CreatePwdReq{
		Name:    req.Name,
		Url:     req.Url,
		Account: req.Account,
		Pwd:     req.Pwd,
		Remark:  req.Remark,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreatePwdResp{
		Code: rpcResp.Code,
		Msg:  rpcResp.Msg,
	}, nil
}
