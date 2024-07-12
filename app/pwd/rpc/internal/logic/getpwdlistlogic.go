package logic

import (
	"context"

	"go-zero-pwd/app/pwd/rpc/internal/svc"
	"go-zero-pwd/app/pwd/rpc/pwd"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPwdListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPwdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPwdListLogic {
	return &GetPwdListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPwdListLogic) GetPwdList(in *pwd.GetPwdListReq) (*pwd.GetPwdListResp, error) {
	// todo: add your logic here and delete this line

	return &pwd.GetPwdListResp{}, nil
}
