package goJwt

import (
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/gin-gonic/gin"
)

const (
	contextTokenKey = "Authentication"
)

//GetJwtStatus 从ctx Keys获取jwt描述
func GetJwtStatus(ctx *gin.Context) (*go_protocol.JwtStatus, error) {
	// 从上下文火获取
	jwtStatusInterface, exist := ctx.Get(contextTokenKey)
	if !exist {
		return nil, go_protocol.LoginInfoNotExistError
	}

	// reflect
	jwtStatus, ok := jwtStatusInterface.(go_protocol.JwtStatus)
	if !ok {
		return nil, go_protocol.LoginInfoNotExistError
	}

	// 判空
	if jwtStatus.GetUid() == "" || jwtStatus.GetType() == "" {
		return nil, go_protocol.LoginInfoEmptyParamError
	}
	return &jwtStatus, nil
}

//GetLoginInfo 从ctx Keys获取登录态
func GetLoginInfo(ctx *gin.Context) (*go_protocol.LoginStatus, error) {

	// 获取jwt描述
	jwtStatus, err := GetJwtStatus(ctx)
	if err != nil {
		return nil, err
	}
	return &jwtStatus.LoginStatus, nil
}
