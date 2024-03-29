package category

import (
	"context"
	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"
	"gincmf/service/portal/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListGetReq) (resp types.Response) {
	c := l.svcCtx
	db := c.Db
	category := model.PortalCategory{
		ParentId: 0,
	}
	data, err := category.ListWithTree(db)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", data)
	return
}
