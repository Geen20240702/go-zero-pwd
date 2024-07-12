package logic

import (
	"context"
	"go-zero-pwd/app/pwd/rpc/internal/svc"
	"go-zero-pwd/app/pwd/rpc/model"
	"go-zero-pwd/app/pwd/rpc/pwd"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePwdLogic {
	return &CreatePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePwdLogic) CreatePwd(in *pwd.CreatePwdReq) (*pwd.CreatePwdResp, error) {
	_, err := l.svcCtx.PwdModel.Insert(l.ctx, &model.Pwd{
		Name:    in.Name,
		Url:     in.Url,
		Account: in.Account,
		Pwd:     in.Pwd,
		Remark:  in.Remark,
	})

	if err != nil {
		return nil, err
	}

	return &pwd.CreatePwdResp{
		Code: "1",
		Msg:  "Success",
	}, nil
}
