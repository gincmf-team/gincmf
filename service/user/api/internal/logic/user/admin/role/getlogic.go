package role

import (
	"context"
	"gincmf/common/bootstrap/data"
	"gincmf/service/user/model"
	"strings"

	"gincmf/service/user/api/internal/svc"
	"gincmf/service/user/api/internal/types"

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

func (l *GetLogic) Get(req *types.RoleGet) (resp *types.Response, err error) {

	resp = new(types.Response)

	c := l.svcCtx
	db := c.Db
	r := c.Request

	// todo: add your logic here and delete this line
	var query []string
	var queryArgs []interface{}

	status := req.Status
	if status != "" {
		query = append(query, "status = ?")
		queryArgs = append(queryArgs, status)
	}
	// 名称模糊搜索
	name := req.Name
	if name != "" {
		query = append(query, "name LIKE ?")
		queryArgs = append(queryArgs, "%"+name+"%")
	}

	var queryStr string
	queryStr = strings.Join(query, " AND ")
	current, pageSize, err := new(data.Paginate).Default(r)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}

	result, err := new(model.Role).Paginate(db, current, pageSize, queryStr, queryArgs)
	if err != nil {
		resp.Error(err.Error(), nil)
		return
	}
	resp.Success("获取成功！", result)
	return
}
