/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 17:09:31
 * @LastEdit: Sy.
 * @FilePath: \server\controllers\admin_controllers\admin_auth_controller.go
 * @Description: 权限因子
 */
package admin_controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"vue-typescript-beego-admin/server/utils"

	"strconv"

	cache "github.com/patrickmn/go-cache"
	"vue-typescript-beego-admin/server/models"
)

type AuthController struct {
	BaseController
}

// 列表
func (_this *AuthController) List() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.AuthGetList(1, 1000, filters...)

	list := GetChildrens(0, result)

	_this.ajaxList(count, list)
}

// 节点详情
func (_this *AuthController) Info() {
	id, _ := _this.GetInt("id")
	result, _ := models.AuthGetById(id)
	// if err == nil {
	// 	_this.ajaxMsg(err.Error(), MSG_ERR)
	// }
	row := make(map[string]interface{})
	row["id"] = result.Id
	row["pid"] = result.Pid
	row["auth_name"] = result.AuthName
	row["auth_url"] = result.AuthUrl
	row["sort"] = result.Sort
	row["is_show"] = result.IsShow
	row["icon"] = result.Icon

	fmt.Println(row)

	_this.ajaxList(0, row)
}

//新增或修改
func (_this *AuthController) Edit() {
	listStr := _this.GetString("list")
	// 列表添加、修改
	_this.EditAuths(listStr)

	// auth := new(models.Auth)
	// auth.UserId = _this.userId
	// auth.Pid, _ = _this.GetInt("pid")
	// auth.AuthName = strings.TrimSpace(_this.GetString("auth_name"))
	// auth.AuthUrl = strings.TrimSpace(_this.GetString("auth_url"))
	// auth.Sort, _ = _this.GetInt("sort")
	// auth.IsShow, _ = _this.GetInt("is_show")
	// auth.Icon = strings.TrimSpace(_this.GetString("icon"))
	// auth.UpdateTime = time.Now().Unix()

	// auth.Status = 1

	// if id == 0 {
	// 	_this.ajaxMsg("找不到", MSG_ERR)

	// 	//新增
	// 	auth.CreateTime = time.Now().Unix()
	// 	auth.CreateId = _this.userId
	// 	auth.UpdateId = _this.userId
	// 	if _, err := models.AuthAdd(auth); err != nil {
	// 		_this.ajaxMsg(err.Error(), MSG_ERR)
	// 	}
	// } else {
	// 	auth.Id = id
	// 	auth.UpdateId = _this.userId
	// 	if err := auth.Update(); err != nil {
	// 		_this.ajaxMsg(err.Error(), MSG_ERR)
	// 	}
	// }
	// utils.Che.Set("menu"+strconv.Itoa(_this.user.Id), nil, cache.DefaultExpiration)
	// _this.ajaxMsg("", MSG_OK)
}

// auths edit by list
func (_this *AuthController) EditAuths(jsonStr string) {
	auths := make([]map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(jsonStr), &auths); err != nil {
		_this.ajaxMsg(err.Error(), MSG_ERR)
	}
	if len(auths) == 0 {
		_this.ajaxMsg("缺少参数", CODE_PARAMS)
	}
	for _, v := range auths {
		id := int(v["id"].(float64))
		auth, err := models.AuthGetById(id)
		if err != nil {
			auth := new(models.Auth)
			auth.CreateId = _this.userId
			auth.CreateTime = time.Now().Unix()
		}
		auth.Status = 1
		auth.AuthName = v["title"].(string)
		auth.AuthUrl = v["authUrl"].(string)
		auth.Pid = int(v["pid"].(float64))
		auth.Sort = int(v["sort"].(float64))
		auth.Icon = v["icon"].(string)
		auth.IsShow = int(v["isShow"].(float64))
		auth.Path = v["path"].(string)
		auth.SidebarHidden = int(v["sidebarHidden"].(float64))
		auth.Breadcrumb = int(v["breadcrumb"].(float64))
		auth.Component = v["component"].(string)
		auth.Redirect = v["redirect"].(string)
		auth.UpdateId = _this.userId
		auth.UpdateTime = time.Now().Unix()
		auth.Update()
	}
	_this.ajaxMsg("", MSG_OK)
}

//删除
func (_this *AuthController) Del() {
	id, _ := _this.GetInt("id")
	auth, _ := models.AuthGetById(id)
	auth.Id = id
	auth.Status = 0
	if err := auth.Update(); err != nil {
		_this.ajaxMsg(err.Error(), MSG_ERR)
	}
	utils.Che.Set("menu"+strconv.Itoa(_this.user.Id), nil, cache.DefaultExpiration)
	_this.ajaxMsg("", MSG_OK)
}

// element tree data
func GetChildrens(pid int, authList []*models.Auth) []map[string]interface{} {
	list := make([]map[string]interface{}, 0)
	for _, v := range authList {
		if v.Pid == pid {
			row := make(map[string]interface{})
			row["id"] = v.Id
			row["title"] = v.AuthName
			row["authUrl"] = v.AuthUrl
			row["isShow"] = v.IsShow
			row["pid"] = v.Pid
			row["sort"] = v.Sort
			row["path"] = v.Path
			row["sidebarHidden"] = v.SidebarHidden
			row["breadcrumb"] = v.Breadcrumb
			row["component"] = v.Component
			row["redirect"] = v.Redirect
			row["icon"] = v.Icon

			subList := GetChildrens(v.Id, authList)
			if len(subList) > 0 {
				row["children"] = subList
			}
			list = append(list, row)
		}
	}
	return list
}
