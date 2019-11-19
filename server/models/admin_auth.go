/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 18:29:47
 * @LastEdit: Sy.
 * @FilePath: \server\models\auth.go
 * @Description: 权限因子 前后端路由控制
 */

package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Auth struct {
	Id            int
	AuthName      string
	AuthUrl       string
	UserId        int
	Pid           int
	Sort          int
	Icon          string
	IsShow        int
	Status        int
	Path          string // 前端路由
	SidebarHidden int    // 前端路由添加;导航栏隐藏 默认0-false;
	Breadcrumb    int    // 显示面包屑;默认1-true;
	Component     string // 前端组件名
	Redirect      string // 前端重定向
	CreateId      int
	UpdateId      int
	CreateTime    int64
	UpdateTime    int64
}

const TABLE_AUTH = "admin_auth"

func (a *Auth) TableName() string {
	return TableName(TABLE_AUTH)
}

func AuthGetList(page, pageSize int, filters ...interface{}) ([]*Auth, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Auth, 0)
	query := orm.NewOrm().QueryTable(TableName(TABLE_AUTH))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("pid", "sort").Limit(pageSize, offset).All(&list)

	return list, total
}

func AuthGetListByIds(authIds string, userId int) ([]*Auth, error) {

	list1 := make([]*Auth, 0)
	var list []orm.Params
	//list:=[]orm.Params
	var err error
	if userId == 1 {
		//超级管理员
		_, err = orm.NewOrm().Raw("select id,auth_name,auth_url,pid,icon,is_show from pp_uc_auth where status=? order by pid asc,sort asc", 1).Values(&list)
	} else {
		_, err = orm.NewOrm().Raw("select id,auth_name,auth_url,pid,icon,is_show from pp_uc_auth where status=1 and id in("+authIds+") order by pid asc,sort asc", authIds).Values(&list)
	}

	for k, v := range list {
		fmt.Println(k, v)
	}

	fmt.Println(list)
	return list1, err
}

func AuthAdd(auth *Auth) (int64, error) {
	return orm.NewOrm().Insert(auth)
}

func AuthGetById(id int) (*Auth, error) {
	a := new(Auth)

	err := orm.NewOrm().QueryTable(TableName(TABLE_AUTH)).Filter("id", id).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Auth) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}
