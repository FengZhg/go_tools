package goJwt

import (
	"github.com/FengZhg/go_tools/go_protocol"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	jwtHeaderKey = "token"
)

//AuthMiddleware 获取jwt身份校验中间件
func (g *goJwt) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g.authMiddleware(ctx)
	}
}

//authMiddleware jwt身份校验中间件
func (g *goJwt) authMiddleware(ctx *gin.Context) {
	//从请求头部获取jwt身份描述
	claimStr := g.getClaimStrFromHeader(ctx)

	// 进行校验
	jwtToken, err := jwt.ParseWithClaims(claimStr, &go_protocol.JwtStatus{}, func(token *jwt.Token) (interface{}, error) {
		return g.publicKey, nil
	})
	if err != nil {
		log.Errorf("User Jwt Claim Verify Error err:%v", err)
		ctx.Error(err)
		ctx.Abort()
	}

	// 解析登录态信息
	jwtStatus := jwtToken.Claims.(*go_protocol.JwtStatus)
	ctx.Set(contextTokenKey, jwtStatus)
	ctx.Next()
}

//getClaimStrFromHeader 从请求头部获取jwt身份描述信息
func (g *goJwt) getClaimStrFromHeader(ctx *gin.Context) string {
	return ctx.GetHeader(jwtHeaderKey)
}
