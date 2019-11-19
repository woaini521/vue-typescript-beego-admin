/*
 * @Author: Sy.
 * @Create: 2019-11-13 22:48:45
 * @LastTime: 2019-11-16 17:08:57
 * @LastEdit: Sy.
 * @FilePath: \server\routers\admin.go
 * @Description: 后台管理路由
 */
package routers

import (
	"github.com/astaxie/beego"
	"vue-typescript-beego-admin/server/controllers/admin_controllers"
)

func init() {
	//初始化 namespace
	ns := beego.NewNamespace("/api/admin",
		beego.NSRouter("login", &admin_controllers.LoginController{}, "*:LoginIn"),
		beego.NSRouter("logout", &admin_controllers.LoginController{}, "*:LoginOut"),

		beego.NSRouter("user/list", &admin_controllers.AdminUserController{}, "*:List"),
		beego.NSRouter("user/info", &admin_controllers.AdminUserController{}, "*:Info"),
		beego.NSRouter("user/edit", &admin_controllers.AdminUserController{}, "*:Edit"),
		beego.NSRouter("user/del", &admin_controllers.AdminUserController{}, "*:Del"),

		beego.NSRouter("auth/list", &admin_controllers.AuthController{}, "*:List"),
		beego.NSRouter("auth/edit", &admin_controllers.AuthController{}, "*:Edit"),

		beego.NSRouter("role/list", &admin_controllers.RoleController{}, "*:List"),
		beego.NSRouter("role/info", &admin_controllers.RoleController{}, "*:Info"),
		beego.NSRouter("role/edit", &admin_controllers.RoleController{}, "*:Edit"),
		beego.NSRouter("role/del", &admin_controllers.RoleController{}, "*:Del"),
	)
	//注册 namespace
	beego.AddNamespace(ns)
}
