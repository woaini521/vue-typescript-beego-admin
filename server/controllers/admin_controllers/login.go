/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-16 17:10:29
 * @LastEdit: Sy.
 * @FilePath: \server\controllers\admin_controllers\login.go
 * @Description: 登录
 */

package admin_controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"vue-typescript-beego-admin/server/libs"
	"vue-typescript-beego-admin/server/models"
	"vue-typescript-beego-admin/server/utils"

	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
)

type LoginController struct {
	BaseController
}

//登录 TODO:XSRF过滤
func (_this *LoginController) LoginIn() {
	if _this.userId > 0 {
		_this.redirect(beego.URLFor("HomeController.Index"))
	}
	beego.ReadFromRequest(&_this.Controller)
	if !_this.isPost() {
		_this.ajaxMsg("不支持的请求方法", MSG_ERR)
	}

	username := strings.TrimSpace(_this.GetString("username"))
	password := strings.TrimSpace(_this.GetString("password"))

	if username != "" && password != "" {
		user, err := models.AdminGetByName(username)
		fmt.Println(user)
		flash := beego.NewFlash()
		errorMsg := ""
		if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
			errorMsg = "帐号或密码错误"
		} else if user.Status == 0 {
			errorMsg = "该帐号已禁用"
		} else {
			user.LastIp = _this.getClientIp()
			user.LastLogin = time.Now().Unix()
			user.Update()
			utils.Che.Set("uid"+strconv.Itoa(user.Id), user, cache.DefaultExpiration)
			authkey := libs.Md5([]byte(_this.getClientIp() + "|" + user.Password + user.Salt))
			_this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)

			// 返回信息
			res := make(map[string]interface{})
			res["auth"] = strconv.Itoa(user.Id) + "|" + authkey
			res["loginName"] = user.RealName
			res["roles"] = user.RoleIds
			_this.ajaxObj(res, "", MSG_OK)
		}
		_this.ajaxMsg(errorMsg, MSG_ERR)
		flash.Error(errorMsg)
		flash.Store(&_this.Controller)
	}
	_this.ajaxMsg("用户名或密码不能为空", CODE_PARAMS)
}

//登出
func (_this *LoginController) LoginOut() {
	_this.Ctx.SetCookie("auth", "")
	utils.Che.Delete("uid" + strconv.Itoa(_this.userId))
	_this.ajaxMsg("", MSG_OK)
}

func (_this *LoginController) NoAuth() {
	_this.ajaxMsg("没有权限", MSG_ERR)
}
