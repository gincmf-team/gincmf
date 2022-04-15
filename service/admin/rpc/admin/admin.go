// Code generated by goctl. DO NOT EDIT!
// Source: admin.proto

package admin

import (
	"context"

	"gincmf/service/admin/rpc/types/admin"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AdminMenu      = admin.AdminMenu
	AdminMenuReply = admin.AdminMenuReply
	AdminMenuReq   = admin.AdminMenuReq

	Admin interface {
		GetMenus(ctx context.Context, in *AdminMenuReq, opts ...grpc.CallOption) (*AdminMenuReply, error)
	}

	defaultAdmin struct {
		cli zrpc.Client
	}
)

func NewAdmin(cli zrpc.Client) Admin {
	return &defaultAdmin{
		cli: cli,
	}
}

func (m *defaultAdmin) GetMenus(ctx context.Context, in *AdminMenuReq, opts ...grpc.CallOption) (*AdminMenuReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.GetMenus(ctx, in, opts...)
}