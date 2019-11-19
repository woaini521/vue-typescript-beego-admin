/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 17:09:35
 * @LastEdit: Sy.
 * @FilePath: \server\controllers\admin_controllers\admin_role_controller.go
 * @Description: 角色
 */

package admin_controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"vue-typescript-beego-admin/server/models"
)

type RoleController struct {
	BaseController
}

func (_this *RoleController) Info() {
	id, _ := _this.GetInt("id", 0)

	//获取选择的树节点
	roleAuth, _ := models.RoleAuthGetById(id)
	authId := make([]int, 0)
	for _, v := range roleAuth {
		authId = append(authId, v.AuthId)
	}
	row := make(map[string]interface{})
	row["authIds"] = authId
	_this.ajaxObj(row, "", MSG_OK)
}

func (_this *RoleController) Edit() {
	role := new(models.Role)
	role.RoleName = strings.TrimSpace(_this.GetString("roleName"))
	role.Detail = strings.TrimSpace(_this.GetString("detail"))
	role.CreateTime = time.Now().Unix()
	role.UpdateTime = time.Now().Unix()
	role.Status = 1
	auths := strings.TrimSpace(_this.GetString("nodesData"))
	role_id, _ := _this.GetInt("id")
	if role_id <= 0 {
		//新增
		role.CreateTime = time.Now().Unix()
		role.UpdateTime = time.Now().Unix()
		role.CreateId = _this.userId
		role.UpdateId = _this.userId
		if id, err := models.RoleAdd(role); err != nil {
			_this.ajaxMsg(err.Error(), MSG_ERR)
		} else {
			ra := new(models.RoleAuth)
			authsSlice := strings.Split(auths, ",")
			for _, v := range authsSlice {
				aid, _ := strconv.Atoi(v)
				ra.AuthId = aid
				ra.RoleId = id
				models.RoleAuthAdd(ra)
			}
		}
		_this.ajaxMsg("", MSG_OK)
	}
	//修改
	role.Id = role_id
	role.UpdateTime = time.Now().Unix()
	role.UpdateId = _this.userId
	if err := role.Update(); err != nil {
		_this.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		// 删除该角色权限
		models.RoleAuthDelete(role_id)
		ra := new(models.RoleAuth)
		authsSlice := strings.Split(auths, ",")
		for _, v := range authsSlice {
			aid, _ := strconv.Atoi(v)
			ra.AuthId = aid
			ra.RoleId = int64(role_id)
			models.RoleAuthAdd(ra)
		}

	}
	_this.ajaxMsg("", MSG_OK)
}

func (_this *RoleController) Del() {

	role_id, _ := _this.GetInt("id")
	role, _ := models.RoleGetById(role_id)
	role.Status = 0
	role.Id = role_id
	role.UpdateTime = time.Now().Unix()

	if err := role.Update(); err != nil {
		_this.ajaxMsg(err.Error(), MSG_ERR)
	}
	// 删除该角色权限
	// models.RoleAuthDelete(role_id)
	_this.ajaxMsg("", MSG_OK)
}

func (_this *RoleController) List() {
	//列表
	page, err := _this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := _this.GetInt("limit")
	if err != nil {
		limit = 30
	}

	roleName := strings.TrimSpace(_this.GetString("roleName"))
	_this.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if roleName != "" {
		filters = append(filters, "role_name__icontains", roleName)
	}
	result, count := models.RoleGetList(page, _this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["roleName"] = v.RoleName
		row["detail"] = v.Detail
		row["createTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["updateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	_this.ajaxList(count, list)
}
