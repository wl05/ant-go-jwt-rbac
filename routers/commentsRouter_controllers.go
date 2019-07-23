package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["ant-go-jwt-rbac/controllers:LoginController"] = append(beego.GlobalControllerRouter["ant-go-jwt-rbac/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-go-jwt-rbac/controllers:LogoutController"] = append(beego.GlobalControllerRouter["ant-go-jwt-rbac/controllers:LogoutController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-go-jwt-rbac/controllers:RegisterController"] = append(beego.GlobalControllerRouter["ant-go-jwt-rbac/controllers:RegisterController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["ant-go-jwt-rbac/controllers:UserController"] = append(beego.GlobalControllerRouter["ant-go-jwt-rbac/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetUsers",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
