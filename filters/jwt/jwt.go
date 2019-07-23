package jwt

import (
	"ant-go-jwt-rbac/common/consts"
	"ant-go-jwt-rbac/common/utils"
	"fmt"
	"github.com/astaxie/beego/context"
	"github.com/go-redis/redis"
	"time"
)

type BaseMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type Data struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type DataMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	} `json:"data"`
}

func Jwt(ctx *context.Context) {

	//fmt.Println(ctx.Request.Header)
	accessToken := ctx.Input.Cookie("accessToken")
	refreshToken := ctx.Input.Cookie("refreshToken")

	isValidAccessToken, _, _ := utils.CheckToken(accessToken)
	isValidRefreshToken, _, refreshTokenClaims := utils.CheckToken(refreshToken)

	// 先判断refreshToken是否已经过期，过期直接返回让客户端重新登录
	if !isValidRefreshToken {
		ctx.Output.JSON(BaseMsg{consts.ERROR_CODE_LOGIN_ERROR, consts.ERROR_DES_LOGIN_ERROR}, false, false)
		return
	}
	// 再判断redis中的黑名单里面是否有isValidRefreshToken如果有说明客户端已经登出，解决jwt登出问题
	val, err := utils.RClient.Get(string(refreshToken)).Result()
	if err != redis.Nil && err != nil {
		ctx.Output.JSON(BaseMsg{consts.ERROR_CODE_REQUEST, consts.ERROR_DES_REQUEST}, false, false)
		return
	}

	if val == "exited" {
		ctx.Output.JSON(BaseMsg{consts.ERROR_CODE_LOGIN_ERROR, consts.ERROR_DES_LOGIN_ERROR}, false, false)
		return
	}
	// refreshToken有效但是accessToken过期，可以重新生成新的accessToken让客户端重新发起请求,解决token刷新问题
	if !isValidAccessToken {
		t, _ := utils.CreateToken(refreshTokenClaims.Email, time.Now().Add(1*time.Minute))
		ctx.Output.Cookie("accessToken", accessToken, "/")   // 设置cookie
		ctx.Output.Cookie("refreshToken", refreshToken, "/") // 设置cookie
		fmt.Println(t)
		ctx.Output.JSON(DataMsg{consts.ERROR_CODE_REFRESH_TOKEN, consts.ERROR_DES_REFRESH_TOKEN, Data{AccessToken: t, RefreshToken: refreshToken}}, false, false)
		return
	}
}
