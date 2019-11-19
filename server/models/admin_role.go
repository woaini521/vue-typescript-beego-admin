/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 18:30:59
 * @LastEdit: Sy.
 * @FilePath: \server\models\role.go
 * @Description: 角色
 */

package models

import (
	"github.com/astaxie/beego/orm"
)

type Role struct {
	Id         int
	RoleName   string
	Detail     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

const TABLE_ROLE = "admin_role"

func (a *Role) TableName() string {
	return TableName(TABLE_ROLE)
}

func RoleGetList(page, pageSize int, filters ...interface{}) ([]*Role, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Role, 0)
	query := orm.NewOrm().QueryTable(TableName(TABLE_ROLE))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func RoleAdd(role *Role) (int64, error) {
	id, err := orm.NewOrm().Insert(role)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func RoleGetById(id int) (*Role, error) {
	r := new(Role)
	err := orm.NewOrm().QueryTable(TableName(TABLE_ROLE)).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Role) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
