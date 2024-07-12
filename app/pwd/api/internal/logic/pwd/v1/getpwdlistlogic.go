package v1

import (
	"context"

	"go-zero-pwd/app/pwd/api/internal/svc"
	"go-zero-pwd/app/pwd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPwdListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPwdListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPwdListLogic {
	return &GetPwdListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPwdListLogic) GetPwdList(req *types.GetPwdListReq) (resp *types.GetPwdListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
