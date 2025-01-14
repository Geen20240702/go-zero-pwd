// Code generated by goctl. DO NOT EDIT.
// Source: pwd.proto

package server

import (
	"context"

	"go-zero-pwd/app/pwd/rpc/internal/logic"
	"go-zero-pwd/app/pwd/rpc/internal/svc"
	"go-zero-pwd/app/pwd/rpc/pwd"
)

type PwdServer struct {
	svcCtx *svc.ServiceContext
	pwd.UnimplementedPwdServer
}

func NewPwdServer(svcCtx *svc.ServiceContext) *PwdServer {
	return &PwdServer{
		svcCtx: svcCtx,
	}
}

func (s *PwdServer) CreatePwd(ctx context.Context, in *pwd.CreatePwdReq) (*pwd.CreatePwdResp, error) {
	l := logic.NewCreatePwdLogic(ctx, s.svcCtx)
	return l.CreatePwd(in)
}

func (s *PwdServer) UpdatePwd(ctx context.Context, in *pwd.UpdatePwdReq) (*pwd.UpdatePwdResp, error) {
	l := logic.NewUpdatePwdLogic(ctx, s.svcCtx)
	return l.UpdatePwd(in)
}

func (s *PwdServer) GetPwdList(ctx context.Context, in *pwd.GetPwdListReq) (*pwd.GetPwdListResp, error) {
	l := logic.NewGetPwdListLogic(ctx, s.svcCtx)
	return l.GetPwdList(in)
}

func (s *PwdServer) DelPwd(ctx context.Context, in *pwd.DelPwdReq) (*pwd.DelPwdResp, error) {
	l := logic.NewDelPwdLogic(ctx, s.svcCtx)
	return l.DelPwd(in)
}
