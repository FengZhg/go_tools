package goJwt

import (
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/gin-gonic/gin"
)

const (
	contextTokenKey = "Authentication"
)

//GetJwtStatus 从ctx Keys获取jwt描述
func GetJwtStatus(ctx *gin.Context) *go_protocol.JwtStatus {
	// 从上下文火获取
	jwtStatusInterface, exist := ctx.Get(contextTokenKey)
	if !exist {
		return nil
	}

	// reflect
	jwtStatus, ok := jwtStatusInterface.(go_protocol.JwtStatus)
	if !ok {
		return nil
	}
	return &jwtStatus
}

//GetLoginInfo 从ctx Keys获取登录态
func GetLoginInfo(ctx *gin.Context) *go_protocol.LoginStatus {
	return GetJwtStatus(ctx).GetLoginStatus()
}
