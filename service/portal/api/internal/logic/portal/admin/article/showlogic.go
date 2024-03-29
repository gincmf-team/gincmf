package article

import (
	"context"
	comModel "gincmf/common/bootstrap/model"
	"gincmf/common/bootstrap/util"
	"gincmf/service/portal/model"
	"strconv"
	"strings"

	"gincmf/service/portal/api/internal/svc"
	"gincmf/service/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.OneReq) (resp types.Response) {

	c := l.svcCtx
	db := c.Db
	id := req.Id

	query := []string{"id = ?", "delete_at = ?"}
	queryStr := strings.Join(query, " AND ")
	queryArgs := []interface{}{id, 0}

	var result struct {
		model.PortalPost
		UserLogin string                 `json:"user_login"`
		Alias     string                 `json:"alias"`
		Keywords  []string               `json:"keywords"`
		Category  []model.PortalCategory `json:"category"`
		Extends   []model.Extends        `json:"extends"`
		Slug      string                 `json:"slug"`
		model.More
	}

	post := model.PortalPost{}

	// 获取当前文章信息
	err := post.Show(db, queryStr, queryArgs)
	if err != nil {
		resp.Error("查询失败："+err.Error(), nil)
		return
	}

	result.PortalPost = post

	if post.PostKeywords != "" {
		result.Keywords = strings.Split(post.PostKeywords, ",")
	}

	// 获取当前文章全部分类
	pQueryArgs := []interface{}{id, 0}
	pc := model.PortalCategory{}
	category, err := pc.ListWithPost(db, "p.id = ? AND p.delete_at = ?", pQueryArgs)
	result.Category = category

	result.Extends = post.MoreJson.Extends

	result.Photos = post.MoreJson.Photos
	result.Files = post.MoreJson.Files

	result.Audio = post.MoreJson.Audio
	result.AudioPrevPath = post.MoreJson.AudioPrevPath

	result.Video = post.MoreJson.Video
	result.VideoPrevPath = post.MoreJson.VideoPrevPath

	fullUrl := "page/" +  strconv.Itoa(id)
	route := comModel.Route{}
	tx := db.Where("full_url", fullUrl).First(&route)

	if util.IsDbErr(tx) != nil {
		resp.Error(tx.Error.Error(), nil)
		return
	}
	result.Alias = route.Url
	resp.Success("获取成功！", result)
	return
}
