/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 18:30:34
 * @LastEdit: Sy.
 * @FilePath: \server\models\role_auth.go
 * @Description: 角色权限控制
 */

package models

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RoleAuth struct {
	AuthId int `orm:"pk"`
	RoleId int64
}

const TABLE_ROLE_AUTH = "admin_role_auth"

func (ra *RoleAuth) TableName() string {
	return TableName(TABLE_ROLE_AUTH)
}

func RoleAuthAdd(ra *RoleAuth) (int64, error) {
	return orm.NewOrm().Insert(ra)
}

func RoleAuthGetById(id int) ([]*RoleAuth, error) {
	list := make([]*RoleAuth, 0)
	query := orm.NewOrm().QueryTable(TableName(TABLE_ROLE_AUTH))
	_, err := query.Filter("role_id", id).All(&list, "AuthId")
	if err != nil {
		return nil, err
	}
	return list, nil
}

func RoleAuthDelete(id int) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName(TABLE_ROLE_AUTH))
	return query.Filter("role_id", id).Delete()
}

//获取多个
func RoleAuthGetByIds(RoleIds string) (Authids string, err error) {
	list := make([]*RoleAuth, 0)
	query := orm.NewOrm().QueryTable(TableName(TABLE_ROLE_AUTH))
	ids := strings.Split(RoleIds, ",")
	_, err = query.Filter("role_id__in", ids).All(&list, "AuthId")
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.Itoa(v.AuthId))
			b.WriteString(",")
		}
	}
	Authids = strings.TrimRight(b.String(), ",")
	return Authids, nil
}

func RoleAuthMultiAdd(ras []*RoleAuth) (n int, err error) {
	query := orm.NewOrm().QueryTable(TableName(TABLE_ROLE_AUTH))
	i, _ := query.PrepareInsert()
	for _, ra := range ras {
		_, err := i.Insert(ra)
		if err == nil {
			n = n + 1
		}
	}
	i.Close() // 别忘记关闭 statement
	return n, err
}
