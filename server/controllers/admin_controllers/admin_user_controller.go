/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 17:09:58
 * @LastEdit: Sy.
 * @FilePath: \server\controllers\admin_controllers\admin_user_controller.go
 * @Description: 管理员 个人中心
 */

package admin_controllers

import (
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"

	"vue-typescript-beego-admin/server/libs"
	"vue-typescript-beego-admin/server/models"
	"vue-typescript-beego-admin/server/utils"
)

type AdminUserController struct {
	BaseController
}

func (_this *AdminUserController) Info() {
	id := _this.userId
	profile, _ := _this.GetInt("profile")
	Admin, _ := models.AdminGetById(id)
	row := make(map[string]interface{})
	row["id"] = Admin.Id
	row["loginName"] = Admin.LoginName
	row["realName"] = Admin.RealName
	row["phone"] = Admin.Phone
	row["email"] = Admin.Email

	if profile == 1 {
		_this.ajaxObj(row, "", MSG_OK)
	}

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if _this.userId != 1 {
		//普通管理员
		adminAuthIds, _ := models.RoleAuthGetByIds(_this.user.RoleIds)
		adminAuthIDArr := strings.Split(adminAuthIds, ",")
		filters = append(filters, "id__in", adminAuthIDArr)
	}
	result, _ := models.AuthGetList(1, 1000, filters...)
	row["list"] = GetNodes(1, result)

	_this.ajaxObj(row, "", MSG_OK)

	// utils.Che.Set("uid"+strconv.Itoa(_this.user.Id), nil, cache.DefaultExpiration)
	// _this.display()
}

func (_this *AdminUserController) List() {
	//列表
	page, err := _this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := _this.GetInt("limit")
	if err != nil {
		limit = 30
	}

	realName := strings.TrimSpace(_this.GetString("realName"))

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	_this.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "delete", 0)
	//
	if realName != "" {
		filters = append(filters, "real_name__icontains", realName)
	}
	result, count := models.AdminGetList(page, _this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["loginName"] = v.LoginName
		row["realName"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["roleIds"] = v.RoleIds
		row["createTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["lastLogin"] = beego.Date(time.Unix(v.LastLogin, 0), "Y-m-d H:i:s")
		row["lastIP"] = v.LastIp
		row["updateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["statusText"] = StatusText[v.Status]
		list[k] = row
	}
	_this.ajaxList(count, list)
}

func (_this *AdminUserController) Edit() {
	AdminID, _ := _this.GetInt("id")
	Admin, _ := models.AdminGetById(AdminID)

	profile, _ := _this.GetBool("profile", false) // 个人中心修改

	if !profile && Admin.Id == 1 {
		_this.ajaxMsg("超级管理员不允许编辑", MSG_ERR)
	}

	editType := _this.GetString("editType")
	if editType == "STATUS" {
		_this.EditStatus(Admin)
	}
	//修改
	Admin.Id = AdminID
	Admin.UpdateTime = time.Now().Unix()
	Admin.UpdateId = _this.userId
	Admin.LoginName = strings.TrimSpace(_this.GetString("loginName"))
	Admin.RealName = strings.TrimSpace(_this.GetString("realName"))
	Admin.Phone = strings.TrimSpace(_this.GetString("phone"))
	Admin.Email = strings.TrimSpace(_this.GetString("email"))

	resetPwd, _ := _this.GetBool("resetPwd", false)

	if resetPwd {

		defer utils.Che.Delete("uid" + strconv.Itoa(AdminID))
		pwd, salt := libs.Password(4, "")
		if profile {
			pwdOld := strings.TrimSpace(_this.GetString("passwordOld"))
			pwdOldMd5 := libs.Md5([]byte(pwdOld + Admin.Salt))
			if Admin.Password != pwdOldMd5 {
				_this.ajaxMsg("旧密码错误", MSG_ERR)
			}

			pwdNew1 := strings.TrimSpace(_this.GetString("password"))
			pwdNew2 := strings.TrimSpace(_this.GetString("passwordSure"))

			if pwdNew1 != pwdNew2 {
				_this.ajaxMsg("两次密码不一致", MSG_ERR)
			}

			pwd, salt = libs.Password(4, pwdNew1)
		}

		Admin.Password = pwd
		Admin.Salt = salt
	}
	Admin.UpdateTime = time.Now().Unix()
	Admin.UpdateId = _this.userId
	Admin.Status = 1

	var err error

	if AdminID == 1 {
		fields := []string{"Password", "Salt", "Phone", "Email", "UpdateTime", "UpdateId"}
		err = Admin.Update(fields...)
	} else {
		err = Admin.Update()
	}

	if err != nil {
		_this.ajaxMsg(err.Error(), MSG_ERR)
	}
	_this.ajaxMsg("", MSG_OK)
}

// 状态
func (_this *AdminUserController) EditStatus(admin *models.Admin) {
	status, _ := _this.GetInt("status", -1)
	if status < 0 || status > 1 {
		_this.ajaxMsg("参数错误", CODE_PARAMS)
	}
	admin.UpdateTime = time.Now().Unix()
	admin.UpdateId = _this.userId
	admin.Status = status
	if err := admin.Update(); err != nil {
		_this.ajaxMsg(err.Error(), MSG_ERR)
	}
	_this.ajaxMsg("", MSG_OK)
}

//删除
func (_this *AdminUserController) Del() {
	id, _ := _this.GetInt("id")
	if id == 1 {
		_this.ajaxMsg("超级管理员不允许删除", MSG_ERR)
	}
	admin, err := models.AdminGetById(id)
	if err != nil {
		_this.ajaxMsg(err.Error(), MSG_OK)
	}
	admin.Delete = 1
	admin.UpdateTime = time.Now().Unix()
	admin.UpdateId = _this.userId
	if err := admin.Update(); err != nil {
		_this.ajaxMsg(err.Error(), MSG_ERR)
	}
	_this.ajaxMsg("", MSG_OK)
}

// 组合当前用户权限列表
func GetNodes(pid int, authList []*models.Auth) []map[string]interface{} {
	list := make([]map[string]interface{}, 0)
	for _, v := range authList {
		if v.Pid == pid && v.IsShow == 1 {
			row := make(map[string]interface{})
			row["id"] = v.Id
			row["title"] = v.AuthName
			row["path"] = v.Path
			row["component"] = v.Component
			row["icon"] = v.Icon
			row["redirect"] = v.Redirect
			subList := GetNodes(v.Id, authList)
			if len(subList) > 0 {
				row["children"] = subList
			}
			list = append(list, row)
		}
	}
	return list
}
