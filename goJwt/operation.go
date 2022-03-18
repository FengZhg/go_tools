package goJwt

import (
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	contextTokenKey = "Authentication"
)

//GetLoginInfo 从ctx Keys获取登录态
func GetLoginInfo(ctx *gin.Context) (*go_protocol.LoginStatus, error) {
	// 从上下文火获取
	loginInfoInterface, exist := ctx.Get(contextTokenKey)
	if !exist {
		return nil, go_protocol.LoginInfoNotExistError
	}

	// reflect
	loginInfo, ok := loginInfoInterface.(go_protocol.LoginStatus)
	if !ok {
		return nil, go_protocol.LoginInfoNotExistError
	}

	// 判空
	if loginInfo.GetUid() == "" || loginInfo.GetType() == "" {
		return nil, go_protocol.LoginInfoEmptyParamError
	}
	return &loginInfo, nil
}

//UpgradeLoginInfo 升级登录态
func UpgradeLoginInfo(ctx *gin.Context, j *goJwt) (string, error) {
	// 获取登录态
	loginInfo, err := GetLoginInfo(ctx)
	if err != nil {
		log.Errorf("Get Login Info Error err:%v", err)
		return "", err
	}

	// 获取新的jwt描述，并且直接返回
	return j.ApplyToken(loginInfo.GetUid())
}
