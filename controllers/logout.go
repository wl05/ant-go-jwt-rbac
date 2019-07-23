package controllers

import (
	"ant-go-jwt-rbac/common/consts"
	"ant-go-jwt-rbac/common/utils"
	"github.com/astaxie/beego"
)

// Operations about Logout
type LogoutController struct {
	beego.Controller
}

// @Title 用户登出
// @Description 用户登出
// @Success 200 登出成功
// @router / [post]
func (this *LogoutController) Logout() {
	refreshToken := this.Ctx.GetCookie("refreshToken")
	utils.RClient.Set(string(refreshToken), "exited", 0) // 这还要设置一下过期时间

	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"msg":  "登出成功",
	}
	this.ServeJSON()
	return
}
