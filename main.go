package main

import (
	"ant-go-jwt-rbac/common/utils"
	"ant-go-jwt-rbac/filters/jwt"
	_ "ant-go-jwt-rbac/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	utils.MysqlClient()
	utils.RedisClient()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, jwt.Jwt)
	//is, _ := casbin.NewEnforcer("./authz/authz_model.conf", "./authz/authz_policy.csv")
	//beego.InsertFilter("*", beego.BeforeRouter, authz.NewAuthorizer(is))
	beego.Run()
}
