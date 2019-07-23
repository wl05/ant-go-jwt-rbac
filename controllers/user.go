package controllers

import (
	"ant-go-jwt-rbac/common/consts"
	"ant-go-jwt-rbac/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @获取用户列表
// @Description 获取用户列表的
// @Success 200 请求成功
// @Success 1102   请求出错
// @router / [get]
func (this *UserController) GetUsers() {
	users, err := models.GetUsers()
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code": consts.ERROR_CODE_REQUEST,
			"msg":  consts.ERROR_DES_REQUEST,
		}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{
		"code": consts.SUCCECC,
		"data": users,
	}
	this.ServeJSON()
	return
}
