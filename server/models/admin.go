/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 18:36:07
 * @LastEdit: Sy.
 * @FilePath: \server\models\admin.go
 * @Description: 管理员
 */

package models

import (
	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	Delete     int
	CreateTime int64
	UpdateTime int64
}

const TABLE_ADMIN = "admin_user"

func (a *Admin) TableName() string {
	return TableName(TABLE_ADMIN)
}

func AdminAdd(a *Admin) (int64, error) {
	return orm.NewOrm().Insert(a)
}

func AdminGetByName(loginName string) (*Admin, error) {
	a := new(Admin)
	err := orm.NewOrm().QueryTable(TableName(TABLE_ADMIN)).Filter("login_name", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func AdminGetList(page, pageSize int, filters ...interface{}) ([]*Admin, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Admin, 0)
	query := orm.NewOrm().QueryTable(TableName(TABLE_ADMIN))
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

func AdminGetById(id int) (*Admin, error) {
	r := new(Admin)
	err := orm.NewOrm().QueryTable(TableName(TABLE_ADMIN)).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

// func RoleAuthDelete(id int) (int64, error) {
// 	query := orm.NewOrm().QueryTable(TableName("role_auth"))
// 	return query.Filter("role_id", id).Delete()
// }

// func RoleAuthMultiAdd(ras []*RoleAuth) (n int, err error) {
// 	query := orm.NewOrm().QueryTable(TableName("role_auth"))
// 	i, _ := query.PrepareInsert()
// 	for _, ra := range ras {
// 		_, err := i.Insert(ra)
// 		if err == nil {
// 			n = n + 1
// 		}
// 	}
// 	i.Close() // 别忘记关闭 statement
// 	return n, err
// }
