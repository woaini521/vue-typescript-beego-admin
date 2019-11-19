/*
 * @Author: Sy.
 * @Create: 2019-11-01 20:54:15
 * @LastTime: 2019-11-18 21:48:58
 * @LastEdit: Sy.
 * @FilePath: \server\controllers\admin_controllers\base_controller.go
 * @Description: 后台管理
 */

package admin_controllers

import (
	"strconv"
	"strings"

	"vue-typescript-beego-admin/server/libs"
	"vue-typescript-beego-admin/server/models"
	"vue-typescript-beego-admin/server/utils"

	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
)

const (
	MSG_OK          = 0
	CODE_LOGIN_EXIT = 5000 // 登录过期
	MSG_ERR         = 1000 // 一般错误
	CODE_PARAMS     = 1001 // 参数错误

)

const DEFAULT_MESSAGE = "成功"

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.Admin
	userId         int
	userName       string
	loginName      string
	pageSize       int
	allowUrl       string
}

//前期准备
func (_this *BaseController) Prepare() {

	_this.pageSize = 20
	controllerName, actionName := _this.GetControllerAndAction()
	_this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	_this.actionName = strings.ToLower(actionName)
	_this.Data["version"] = beego.AppConfig.String("version")
	_this.Data["siteName"] = beego.AppConfig.String("site.name")
	_this.Data["curRoute"] = _this.controllerName + "." + _this.actionName
	_this.Data["curController"] = _this.controllerName
	_this.Data["curAction"] = _this.actionName

	_this.auth()

	_this.Data["loginUserId"] = _this.userId
	_this.Data["loginUserName"] = _this.userName
}

//登录权限验证
func (_this *BaseController) auth() {
	// arr := strings.Split(_this.Ctx.GetCookie("auth"), "|")
	token := _this.Ctx.Input.Header("X-Access-Token")
	arr := strings.Split(token, "|")

	_this.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			var err error

			cheUser, found := utils.Che.Get("uid" + strconv.Itoa(userId))
			user := &models.Admin{}
			if found && cheUser != nil { //从缓存取用户
				user = cheUser.(*models.Admin)
			} else {
				user, err = models.AdminGetById(userId)
				utils.Che.Set("uid"+strconv.Itoa(userId), user, cache.DefaultExpiration)
			}
			if password != libs.Md5([]byte(_this.getClientIp()+"|"+user.Password+user.Salt)) {
				_this.ajaxMsg("认证失效,请重新登录", CODE_LOGIN_EXIT)
			}
			if err == nil {
				_this.userId = user.Id

				_this.loginName = user.LoginName
				_this.userName = user.RealName
				_this.user = user

				_this.AdminAuth()
			}

			isHasAuth := strings.Contains(_this.allowUrl, _this.controllerName+"/"+_this.actionName)

			//不需要权限检查
			noAuth := ""

			isNoAuth := strings.Contains(noAuth, _this.actionName)

			if !isHasAuth && !isNoAuth {
				beego.Info(_this.controllerName + "/" + _this.actionName)
				_this.ajaxMsg("没有权限", MSG_ERR)
			}
		}
	}

	if _this.userId == 0 && (_this.controllerName != "login" && _this.actionName != "loginin") {
		_this.redirect(beego.URLFor("LoginController.LoginIn"))
	}
}

func (_this *BaseController) AdminAuth() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if _this.userId != 1 {
		//普通管理员
		adminAuthIds, _ := models.RoleAuthGetByIds(_this.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters := make([]interface{}, 0)
		filters = append(filters, "id__in", adminAuthIdArr)
	}
	result, _ := models.AuthGetList(1, 1000, filters...)
	allowUrl := ""
	for _, v := range result {
		v.AuthUrl = strings.TrimSpace(v.AuthUrl)
		if v.AuthUrl != "" && v.AuthUrl != "/" {
			allowUrl += v.AuthUrl
		}
	}
	_this.allowUrl = allowUrl
}

// 是否POST提交
func (_this *BaseController) isPost() bool {
	return _this.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (_this *BaseController) getClientIp() string {
	s := _this.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

// 重定向
func (_this *BaseController) redirect(url string) {
	_this.Redirect(url, 302)
	_this.StopRun()
}

//加载模板
func (_this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = _this.controllerName + "/" + _this.actionName + ".html"
	}
	_this.Layout = "public/layout.html"
	_this.TplName = tplname
}

//ajax返回
func (_this *BaseController) ajaxMsg(msg interface{}, code int) {
	out := make(map[string]interface{})
	out["code"] = code
	out["message"] = msg
	out["data"] = nil
	_this.Data["json"] = out
	_this.ServeJSON()
	_this.StopRun()
}

func (_this *BaseController) ajaxObj(data interface{}, msg string, code int) {
	out := make(map[string]interface{})
	out["code"] = code
	out["message"] = msg
	out["data"] = data
	_this.Data["json"] = out
	_this.ServeJSON()
	_this.StopRun()
}

//ajax返回 列表
func (_this *BaseController) ajaxList(total int64, list interface{}) {
	out := make(map[string]interface{})
	out["code"] = 0
	out["message"] = DEFAULT_MESSAGE

	data := make(map[string]interface{})
	data["total"] = total
	data["list"] = list

	out["data"] = data

	// out["count"] = count
	// out["data"] = data
	_this.Data["json"] = out
	_this.ServeJSON()
	_this.StopRun()
}
