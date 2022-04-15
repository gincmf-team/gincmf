// Code generated by goctl. DO NOT EDIT!
// Source: admin.proto

package server

import (
	"context"

	"gincmf/service/admin/rpc/internal/logic"
	"gincmf/service/admin/rpc/internal/svc"
	"gincmf/service/admin/rpc/types/admin"
)

type AdminServer struct {
	svcCtx *svc.ServiceContext
	admin.UnimplementedAdminServer
}

func NewAdminServer(svcCtx *svc.ServiceContext) *AdminServer {
	return &AdminServer{
		svcCtx: svcCtx,
	}
}

func (s *AdminServer) GetMenus(ctx context.Context, in *admin.AdminMenuReq) (*admin.AdminMenuReply, error) {
	l := logic.NewGetMenusLogic(ctx, s.svcCtx)
	return l.GetMenus(in)
}
