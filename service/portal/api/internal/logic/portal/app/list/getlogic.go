package list

import (
	"context"
	"gincmf/common/bootstrap/data"
	"gincmf/service/portal/model"
	"strings"

	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.PostListReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db

	r := c.Request

	extra := map[string]string{}

	hot := req.Hot

	if hot == 1 {
		extra["hot"] = "1"
	}

	var query []string
	var queryArgs []interface{}

	for _, v := range req.Ids {
		query = append(query, "cp.category_id = ?")
		queryArgs = append(queryArgs, v)
	}

	queryRes := []string{"p.post_type = 1 AND p.delete_at = 0"}
	if len(query) > 0 {
		queryStr := strings.Join(query, " OR ")
		queryRes = append(queryRes, queryStr)
	}

	current, pageSize, err := new(data.Paginate).Default(r)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	data, err := new(model.PortalPost).IndexByCategory(db, current, pageSize, strings.Join(queryRes, " AND "), queryArgs, extra)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	resp.Success("获取成功！", data)
	return
}
