/**
** @创建时间: 2021/11/26 10:48
** @作者　　: return
** @描述　　:
 */

package model

import (
	"gincmf/common/bootstrap/paginate"
	"gincmf/common/bootstrap/util"
	"gorm.io/gorm"
	"time"
)

type Role struct {
	Id        int     `json:"id"`
	ParentId  int     `gorm:"type:int(11);comment:'所属父类id';default:0" json:"parent_id"`
	Name      string  `gorm:"type:varchar(30);comment:'名称'" json:"name"`
	Remark    string  `gorm:"type:varchar(255);comment:'备注'" json:"remark"`
	ListOrder float64 `gorm:"type:float;comment:'排序';default:10000" json:"list_order"`
	CreateAt  int64   `gorm:"type:int(11)" json:"create_at"`
	UpdateAt  int64   `gorm:"type:int(11)" json:"update_at"`
	Status    int     `gorm:"type:tinyint(3);comment:'状态';default:1" json:"status"`
}

type RoleUser struct {
	Id     int `json:"id"`
	RoleId int `gorm:"type:int(11);comment:'角色id';not null" json:"role_id"`
	UserId int `gorm:"type:int(11);comment:'所属用户id';not null" json:"user_id"`
}

func (_ *Role) AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&RoleUser{})

	role := []Role{
		{
			Name:      "超级管理员",
			Remark:    "拥有网站最高管理员权限！",
			ListOrder: 0,
			CreateAt:  time.Now().Unix(),
			UpdateAt:  time.Now().Unix(),
		},
		{
			Name:      "普通管理员",
			Remark:    "权限由最高管理员分配！",
			ListOrder: 0,
			CreateAt:  time.Now().Unix(),
			UpdateAt:  time.Now().Unix(),
		},
	}

	for _, v := range role {
		db.Where("name", v.Name).FirstOrCreate(&v)
	}

}

func (model *Role) Paginate(db *gorm.DB, current, pageSize int, query string, queryArgs []interface{}) (result paginate.Paginate, err error) {

	var role []Role
	var total int64 = 0
	tx := db.Where(query, queryArgs...).Find(&role).Count(&total)
	if tx.Error != nil {
		err = tx.Error
		return
	}
	tx = db.Limit(pageSize).Where(query, queryArgs...).Offset((current - 1) * pageSize).Find(&role)
	if util.IsDbErr(tx) != nil {
		err = tx.Error
		return
	}

	result = paginate.Paginate{Data: role, Current: current, PageSize: pageSize, Total: total}
	if len(role) == 0 {
		result.Data = make([]string, 0)
	}
	return result,nil

}

func (model *Role) Show(db *gorm.DB,query string, queryArgs []interface{}) error {
	tx := db.Where(query,queryArgs...).First(&model)
	if util.IsDbErr(tx) != nil {
		return tx.Error
	}
	return nil
}